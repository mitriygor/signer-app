import {useEffect, useState} from 'react';
import {Card, Col, Row} from "antd";

const WebSocketListener = () => {
    const [signerReqCount, setSignerReqCount] = useState(0);
    const [signerErrorCount, setSignerErrorCount] = useState(0);
    const [brokerReqCount, setBrokerReqCount] = useState(0);
    const [brokerErrorCount, setBrokerErrorCount] = useState(0);
    const [listenerReqCount, setListenerReqCount] = useState(0);
    const [listenerErrorCount, setListenerErrorCount] = useState(0);
    const [loggerReqCount, setLoggerReqCount] = useState(0);
    const [loggerErrorCount, setLoggerErrorCount] = useState(0);

    const [keyKeeperReqCount, setKeyKeeperReqCount] = useState(0);
    const [keyKeeperErrorCount, setKeyKeeperErrorCount] = useState(0);


    const [ws, setWs] = useState(null);


    // Create WebSocket connection on component mount
    useEffect(() => {

        try {
            const ws = new WebSocket('ws://localhost:5002/ws');
            // @ts-ignore
            setWs(ws);

            return () => {
                ws.close();
            };
        } catch (e) {
            console.log("ISSUE", e);
        }

    }, []);

    // Attach message handler
    useEffect(() => {
        if (!ws) return;

        // @ts-ignoreK
        ws.onmessage = (event) => {
            const data = JSON.parse(event.data);
            setSignerReqCount(data.signer_count || 0);
            setSignerErrorCount(data.signer_error_count || 0);
            setBrokerReqCount(data.broker_count || 0);
            setBrokerErrorCount(data.broker_error_count || 0);
            setListenerReqCount(data.listener_count || 0);
            setListenerErrorCount(data.listener_error_count || 0);
            setLoggerReqCount(data.logger_count || 0);
            setLoggerErrorCount(data.logger_error_count || 0);
            setKeyKeeperReqCount(data.key_keeper_count || 0);
            setKeyKeeperErrorCount(data.key_keeper_error_count || 0);
        };

    }, [ws]);

    return (
        <div style={{width: '1200px', margin: '0 auto'}}>
            <Row gutter={16}>
                <Col span={5}>
                    <Card title={'Broker'}>
                        <h1>{brokerReqCount}</h1>
                        <h5>Errors: {brokerErrorCount}</h5>
                    </Card>
                </Col>
                <Col span={5}>
                    <Card title={'Listener'}>
                        <h1>{listenerReqCount}</h1>
                        <h5>Errors: {listenerErrorCount}</h5>
                    </Card>
                </Col>
                <Col span={4}>
                    <Card title={'Key Keeper'}>
                        <h1>{keyKeeperReqCount}</h1>
                        <h5>Errors: {keyKeeperErrorCount}</h5>
                    </Card>
                </Col>
                <Col span={5}>
                    <Card title={'Signer'}>
                        <h1>{signerReqCount}</h1>
                        <h5>Errors: {signerErrorCount}</h5>
                    </Card>
                </Col>
                <Col span={5}>
                    <Card title={'Logger'}>
                        <h1>{loggerReqCount}</h1>
                        <h5>Errors: {loggerErrorCount}</h5>
                    </Card>
                </Col>
            </Row>
        </div>
    );
};

export default WebSocketListener;
