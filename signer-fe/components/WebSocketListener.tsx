import {useEffect, useState} from 'react';
import {Card, Col, Row} from "antd";


const WebSocketListener = ({
                               // @ts-ignore
                               setIsBrokerSubmitting,
                               // @ts-ignore
                               intervalBrokerRef,
                               // @ts-ignore
                               brokerTotal,
                               // @ts-ignore
                               setIsListenerSubmitting,
                               // @ts-ignore
                               intervalListenerRef,
                               // @ts-ignore
                               listenerTotal,
                               // @ts-ignore
                               setIsKeySubmitting,
                               // @ts-ignore
                               intervalKeyRef,
                               // @ts-ignore
                               keyTotal,
                               // @ts-ignore
                               setIsSignerSubmitting,
                               // @ts-ignore
                               intervalSignerRef,
                               // @ts-ignore
                               signerTotal,
                               // @ts-ignore
                               setIsSubmitting,
                               // @ts-ignore
                               intervalRef,
                               // @ts-ignore
                               logsTotal
                           }) => {
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

    const [isBrokerDone, setIsBrokerDone] = useState(false);
    const [isListenerDone, setIsListenerDone] = useState(false);
    const [isKeyDone, setIsKeyDone] = useState(false);
    const [isSignerDone, setIsSignerDone] = useState(false);
    const [isDone, setIsDone] = useState(false);


    // @ts-ignore
    const isDoneStreaming = (total, count, errorCount) => {
        return total != -1 && count + errorCount >= total
    }

    useEffect(() => {

        if (logsTotal === -1) return;
        try {
            const ws = new WebSocket('ws://localhost:5002/ws');
            // @ts-ignore
            setWs(ws);

            setIsBrokerDone(false);
            setIsListenerDone(false);
            setIsKeyDone(false);
            setIsSignerDone(false);
            setIsDone(false);

            return () => {
                ws.close();
            };
        } catch (e) {
            console.log("ISSUE", e);
        }

    }, [logsTotal]);

    useEffect(() => {
        if (!ws) return;

        // @ts-ignoreK
        ws.onmessage = (event) => {
            const data = JSON.parse(event.data);

            const brokerCount = data.broker_count || 0;
            const brokerErrorCount = data.broker_error_count || 0;

            const listenerCount = data.listener_count || 0;
            const listenerErrorCount = data.listener_error_count || 0;

            const keyCount = data.key_keeper_count || 0;
            const keyErrorCount = data.key_keeper_error_count || 0;


            const signerCount = data.signer_count || 0;
            const signerErrorCount = data.signer_error_count || 0;

            const loggerCount = data.logger_count || 0;
            const loggerErrorCount = data.logger_error_count || 0;

            setBrokerReqCount(brokerCount);
            setBrokerErrorCount(brokerErrorCount);

            setListenerReqCount(listenerCount);
            setListenerErrorCount(listenerErrorCount);

            setKeyKeeperReqCount(keyCount);
            setKeyKeeperErrorCount(keyErrorCount);

            setSignerReqCount(signerCount);
            setSignerErrorCount(signerErrorCount);

            setLoggerReqCount(loggerCount);
            setLoggerErrorCount(loggerErrorCount);

            if (!isBrokerDone && isDoneStreaming(brokerTotal, brokerCount, brokerErrorCount)) {
                setIsBrokerSubmitting(false);
                clearInterval(intervalBrokerRef.current);
                setIsBrokerDone(true);
            }

            if (!isListenerDone && isDoneStreaming(listenerTotal, listenerCount, listenerErrorCount)) {
                console.log('DONE LISTENER')
                setIsSubmitting(false);
                clearInterval(intervalRef.current);
                setIsDone(true);
                setIsListenerSubmitting(false);
                clearInterval(intervalListenerRef.current);
                setIsListenerDone(true);
                // @ts-ignore
                ws.close();
            }

            if (!isKeyDone && isDoneStreaming(keyTotal, keyCount, keyErrorCount)) {
                setIsKeySubmitting(false);
                clearInterval(intervalKeyRef.current);
                setIsKeyDone(true);
            }

            if (!isSignerDone && isDoneStreaming(signerTotal, signerCount, signerErrorCount)) {
                setIsSignerSubmitting(false);
                clearInterval(intervalSignerRef.current);
                setIsSignerDone(true);
            }

            if (!isDone && isDoneStreaming(logsTotal, loggerCount, loggerErrorCount)) {
                console.log('DONE LOG')
                setIsSubmitting(false);
                clearInterval(intervalRef.current);
                setIsDone(true);
                setIsListenerSubmitting(false);
                clearInterval(intervalListenerRef.current);
                setIsListenerDone(true);
                // @ts-ignore
                ws.close();
            }
        };
    }, [ws, brokerTotal, listenerTotal, keyTotal, signerTotal, logsTotal]);


    return (
        <div style={{width: '1400px', margin: '0 auto'}}>
            <Row gutter={16}>
                <Col span={5}>
                    <Card title={'Broker'}>
                        <Row align="middle"><h1>{brokerReqCount}</h1><h4>&nbsp; messages</h4></Row>
                        <h5>Errors: {brokerErrorCount}</h5>
                    </Card>
                </Col>
                <Col span={5}>
                    <Card title={'Listener'}>
                        <Row align="middle"><h1>{listenerReqCount}</h1><h4>&nbsp; messages</h4></Row>
                        <h5>Errors: {listenerErrorCount}</h5>
                    </Card>
                </Col>
                <Col span={4}>
                    <Card title={'Key Keeper'}>
                        <Row align="middle"><h1>{keyKeeperReqCount}</h1><h4>&nbsp; keys</h4></Row>
                        <h5>Errors: {keyKeeperErrorCount}</h5>
                    </Card>
                </Col>
                <Col span={5}>
                    <Card title={'Signer'}>
                        <Row align="middle"><h1>{signerReqCount}</h1><h4>&nbsp; signed</h4></Row>
                        <h5>Errors: {signerErrorCount}</h5>
                    </Card>
                </Col>
                <Col span={5}>
                    <Card title={'Logger'}>
                        <Row align="middle"><h1>{loggerReqCount} </h1><h4>&nbsp; logged</h4></Row>
                        <h5>Errors: {loggerErrorCount}</h5>
                    </Card>
                </Col>
            </Row>
        </div>
    );
};

export default WebSocketListener;
