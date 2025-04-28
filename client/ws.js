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

function connect() {
    if (ws && ws.readyState === WebSocket.OPEN) {
        log('已经连接到服务器');
        return;
    }

    const server = document.getElementById('server').value;
    const appId = document.getElementById('app').value;
    const clientId = document.getElementById('client').value;

    ws = new WebSocket(server);

    ws.onopen = () => {
        log('连接已建立');
        updateStatus('已连接');
        document.getElementById('disconnectBtn').disabled = false;

        // 发送认证消息
        const authMsg = JSON.stringify({
            app_id: appId,
            client_id: clientId
        });
        ws.send(authMsg);
        log('已发送认证信息');
    };

    ws.onmessage = (event) => {
        const message = event.data;

        // 处理心跳消息
        if (message === 'heartbeat') {
            log('收到心跳，发送响应...');
            ws.send('heartbeat_ack');
            return;
        }

        log(`收到消息: ${message}`);
    };

    ws.onerror = (error) => {
        log(`连接错误: ${error}`);
        updateStatus('连接错误');
    };

    ws.onclose = () => {
        log('连接已关闭');
        updateStatus('未连接');
        document.getElementById('disconnectBtn').disabled = true;
        if (heartbeatTimer) clearInterval(heartbeatTimer);
    };

    // 启动心跳检测（可选）
    heartbeatTimer = setInterval(() => {
        if (ws.readyState === WebSocket.OPEN) {
            ws.send('heartbeat_check');
        }
    }, 30000);
}

function disconnect() {
    if (ws) {
        ws.close();
        ws = null;
    }
}