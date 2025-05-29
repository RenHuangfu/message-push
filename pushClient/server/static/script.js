const appIdInput = document.getElementById('app');
const clientIdInput = document.getElementById('client');
const connectBtn = document.getElementById('connectBtn');
const disconnectBtn = document.getElementById('disconnectBtn');
const messagesDiv = document.getElementById('messages');
const statusDiv = document.getElementById('status');

let eventSource = null; // 用于 Server-Sent Events

// 在消息区域记录日志
function logMessage(message, type = 'info') {
    const entry = document.createElement('div');
    entry.className = 'message-entry';
    const timestamp = new Date().toLocaleTimeString();
    entry.textContent = `[${timestamp}] ${message}`;
    if (type === 'error') {
        entry.style.color = 'red';
    } else if (type === 'status') {
        entry.style.color = 'blue';
        entry.style.fontStyle = 'italic';
    }
    messagesDiv.appendChild(entry);
    messagesDiv.scrollTop = messagesDiv.scrollHeight; // 自动滚动到底部
}

// 更新页面上的状态显示，并控制按钮的可用性
function updateStatus(newStatus) {
    statusDiv.textContent = `状态: ${newStatus}`;
    // logMessage(`状态变更: ${newStatus}`, 'status'); // 状态变更也会通过SSE的status事件记录，避免重复

    const lowerStatus = newStatus.toLowerCase();
    if (lowerStatus.includes('已连接') || lowerStatus.includes('connected')) {
        connectBtn.disabled = true;
        disconnectBtn.disabled = false;
        appIdInput.disabled = true;
        clientIdInput.disabled = true;
    } else if (lowerStatus.includes('正在连接') || lowerStatus.includes('connecting') || lowerStatus.includes('启动') || lowerStatus.includes('initiating')) {
        connectBtn.disabled = true;
        disconnectBtn.disabled = true; // 连接过程中不允许断开
    } else { // 未连接, 断开, 错误等
        connectBtn.disabled = false;
        disconnectBtn.disabled = true;
        appIdInput.disabled = false;
        clientIdInput.disabled = false;
    }
}

// 点击“连接”按钮时调用
async function connectToServer() {
    const appId = appIdInput.value;
    const clientId = clientIdInput.value;

    if (!appId || !clientId) {
        logMessage('App ID 和 Client ID 不能为空。', 'error');
        return;
    }

    logMessage('尝试通过服务器连接...');
    updateStatus('正在启动连接过程...'); // 初始状态更新
    connectBtn.disabled = true;

    try {
        const response = await fetch('/api/connect', { // 调用后端API
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ appId, clientId })
        });

        const result = await response.json();

        if (!response.ok) {
            logMessage(`启动连接失败: ${result.message || response.statusText}`, 'error');
            updateStatus(`失败: ${result.message || response.statusText}`); // 更新状态为失败
            // connectBtn.disabled = false; // 允许重试，但状态更新应处理此逻辑
            return;
        }

        logMessage(result.message || '服务器已开始连接过程。');
        // 此时，连接状态将由 SSE 更新
        initEventSource(); // 确保 SSE 已初始化或重新初始化

    } catch (error) {
        logMessage(`调用连接 API 出错: ${error.message}`, 'error');
        updateStatus(`API 错误: ${error.message}`);
        // connectBtn.disabled = false;
    }
}

// 点击“断开”按钮时调用
async function disconnectFromServer() {
    logMessage('尝试通过服务器断开连接...');
    disconnectBtn.disabled = true; // 禁用断开按钮，防止重复点击

    try {
        const response = await fetch('/api/disconnect', { // 调用后端API
            method: 'POST',
        });
        const result = await response.json();

        if (!response.ok) {
            logMessage(`启动断开失败: ${result.message || response.statusText}`, 'error');
            updateStatus(`断开失败: ${result.message || response.statusText}`);
            // disconnectBtn.disabled = false; // 如果服务器说它已经断开，可能需要重新启用
        } else {
            logMessage(result.message || '服务器已开始断开过程。');
            // 实际的状态更新将通过 SSE 在客户端确认断开后进行
        }
        // 如果断开成功，SSE 流可能会关闭，或者状态会更新为“已断开”
        // updateStatus("断开已启动"); // 或者等待 SSE 的状态更新

    } catch (error) {
        logMessage(`调用断开 API 出错: ${error.message}`, 'error');
        updateStatus(`API 错误: ${error.message}`);
        // disconnectBtn.disabled = false; // 错误时允许重试
    }
}

// 初始化 Server-Sent Events 连接
function initEventSource() {
    if (eventSource && eventSource.readyState !== EventSource.CLOSED) {
        logMessage('SSE 已连接，无需重新初始化。');
        return;
    }
    if (eventSource) {
        eventSource.close(); // 关闭旧的连接以防万一
    }


    logMessage('正在初始化 Server-Sent Events (SSE) 事件流...');
    eventSource = new EventSource('/api/events'); // 连接到后端的 SSE 端点

    eventSource.onopen = () => {
        logMessage('SSE 连接已建立。');
    };

    // 监听名为 'message' 的事件 (来自 SDK 的 MessageChan)
    eventSource.addEventListener('message', (event) => {
        const messageData = JSON.parse(event.data); // 后端发送的是 JSON 编码的字符串
        logMessage(`收到消息: ${messageData}`);
    });

    // 监听名为 'status' 的事件 (来自 SDK 的 StatusChan)
    eventSource.addEventListener('status', (event) => {
        const statusData = JSON.parse(event.data);
        logMessage(`状态更新: ${statusData}`, 'status');
        updateStatus(statusData); // 更新页面主状态
    });

    // 监听名为 'error' 的事件 (来自 SDK 的 ErrorChan)
    eventSource.addEventListener('error', (event) => {
        const errorData = JSON.parse(event.data);
        logMessage(`错误: ${errorData}`, 'error');
        // updateStatus(`错误: ${errorData}`); // 也可以选择更新主状态显示
    });

    eventSource.onerror = (error) => {
        logMessage('SSE 连接错误或已关闭。', 'error');
        console.error('EventSource 失败:', error);
        // 如果SSE连接失败，可能意味着服务器暂时不可用，或者客户端已经断开。
        // 此时应该允许用户尝试重新连接。
        if (eventSource.readyState === EventSource.CLOSED) {
            updateStatus('SSE 连接已关闭。请尝试重新连接。');
        } else {
            updateStatus('SSE 连接错误。');
        }
        // eventSource.close(); // 它可能会自动关闭
        // connectBtn.disabled = false; // 允许用户尝试重新连接
        // disconnectBtn.disabled = true;
    };
}

// 页面加载时执行
window.onload = () => {
    // 从服务端渲染的HTML中获取初始状态
    const initialStatusText = statusDiv.textContent.replace('状态: ', '').trim();
    updateStatus(initialStatusText); // 根据初始状态设置按钮等

    // 始终尝试初始化SSE，以便获取实时更新。
    // 如果服务器端已有连接，SSE会开始推送其状态。
    initEventSource();
};