let ws = null;
let heartbeatTimer = null;

function log(message) {
    const messagesDiv = document.getElementById('messages');
    const entry = document.createElement('div');
    entry.textContent = `[${new Date().toLocaleTimeString()}] ${message}`;
    messagesDiv.appendChild(entry);
    messagesDiv.scrollTop = messagesDiv.scrollHeight;
}

function updateStatus(status) {
    document.getElementById('status').textContent = status;
}

async function connect() {
    if (ws && ws.readyState === WebSocket.OPEN) {
        log('已经连接到服务器');
        return;
    }

    const serverApiUrl = "http://39.107.58.199:8008/search";
    const appId = document.getElementById('app').value;
    const clientId = document.getElementById('client').value;

    log(`尝试从 ${serverApiUrl} 获取WebSocket连接地址...`);
    updateStatus('正在获取WebSocket地址...');

    try {
        const response = await fetch(serverApiUrl, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                app_id: appId,
                client_id: clientId
            })
        });

        if (!response.ok) {
            const errorText = await response.text();
            log(`获取WebSocket URL失败: ${response.status} ${response.statusText}. 响应: ${errorText}`);
            updateStatus(`获取URL失败: ${response.status}`);
            return;
        }

        const data = await response.json();
        const websocketUrl = data.url;

        if (!websocketUrl) {
            log('从API获取的WebSocket URL无效或不存在。');
            updateStatus('获取的URL无效');
            return;
        }

        log(`成功获取WebSocket URL: ${websocketUrl}`);
        log('开始建立WebSocket连接...');
        updateStatus('正在连接WebSocket...');

        ws = new WebSocket(websocketUrl);

        ws.onopen = () => {
            log(`WebSocket连接已建立: ${websocketUrl}`);
            updateStatus('已连接');
            document.getElementById('disconnectBtn').disabled = false;

            // 发送认证消息
            const authMsg = JSON.stringify({
                app_id: appId,
                client_id: clientId
            });
            ws.send(authMsg);
            log('已发送认证信息到WebSocket');

            // 启动心跳检测（可选）
            if (heartbeatTimer) clearInterval(heartbeatTimer); // 清除旧的计时器
            heartbeatTimer = setInterval(() => {
                if (ws && ws.readyState === WebSocket.OPEN) {
                    ws.send('heartbeat_check');
                    // log('发送心跳检测 (heartbeat_check)'); // 可以取消注释以查看心跳发送日志
                }
            }, 30000);
        };

        ws.onmessage = (event) => {
            const message = event.data;

            // 处理心跳消息
            if (message === 'heartbeat') {
                log('收到心跳 (heartbeat)，发送响应 (heartbeat_ack)...');
                ws.send('heartbeat_ack');
                return;
            }
            // 也可以处理服务器主动发来的 heartbeat_ack (如果服务器会主动发)
            if (message === 'heartbeat_ack') {
                log('收到心跳响应 (heartbeat_ack)');
                return;
            }

            log(`收到消息: ${message}`);
        };

        ws.onerror = (error) => {
            // WebSocket的error事件通常不包含详细的错误信息，更多信息在console
            console.error('WebSocket Error:', error);
            log(`WebSocket连接错误。查看控制台获取更多信息。`);
            updateStatus('WebSocket连接错误');
        };

        ws.onclose = (event) => {
            log(`WebSocket连接已关闭。 Code: ${event.code}, Reason: ${event.reason || 'N/A'}`);
            updateStatus('未连接');
            document.getElementById('disconnectBtn').disabled = true;
            if (heartbeatTimer) clearInterval(heartbeatTimer);
            ws = null; // 确保ws被重置
        };

    } catch (error) {
        console.error('POST请求错误:', error);
        log(`API请求错误: ${error.message}`);
        updateStatus('API请求错误');
    }
}

function disconnect() {
    if (ws) {
        log('手动断开连接...');
        ws.close();
        // onclose事件会自动处理后续状态更新和计时器清理
    } else {
        log('当前没有活动的连接。');
    }
}