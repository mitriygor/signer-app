'use client';
import {Button, Card, Col, ConfigProvider, Form, InputNumber, Row, Switch, theme} from 'antd';
// @ts-ignore
import WebSocketListener from "../../components/WebSocketListener";
// @ts-ignore
import {useEffect, useRef, useState} from "react";


export default function Home() {

    const [isImageOne, setImage] = useState(true);

    const toggleImage = () => {
        setImage(!isImageOne);
    };

    const [form] = Form.useForm();

    // Broker Timer
    const [brokerTimer, setBrokerTimer] = useState(0);
    const [isBrokerSubmitting, setIsBrokerSubmitting] = useState(false);
    const intervalBrokerRef = useRef();
    const [brokerTotal, setBrokerTotal] = useState(-1);


    // Listener Timer
    const [listenerTimer, setListenerTimer] = useState(0);
    const [isListenerSubmitting, setIsListenerSubmitting] = useState(false);
    const intervalListenerRef = useRef();
    const [listenerTotal, setListenerTotal] = useState(-1);


    // Key Timer
    const [keyTimer, setKeyTimer] = useState(0);
    const [isKeySubmitting, setIsKeySubmitting] = useState(false);
    const intervalKeyRef = useRef();
    const [keyTotal, setKeyTotal] = useState(-1);

    // Signer Timer
    const [signerTimer, setSignerTimer] = useState(0);
    const [isSignerSubmitting, setIsSignerSubmitting] = useState(false);
    const intervalSignerRef = useRef();
    const [signerTotal, setSignerTotal] = useState(-1);

    // Logger Timer
    const [timer, setTimer] = useState(0);
    const [isSubmitting, setIsSubmitting] = useState(false);
    const intervalRef = useRef();
    const [logsTotal, setLogsTotal] = useState(-1);


    const handleSubmit = async (values: any) => {
        setIsBrokerSubmitting(true);
        setIsListenerSubmitting(true);
        setIsKeySubmitting(true);
        setIsSignerSubmitting(true);
        setIsSubmitting(true);

        setBrokerTimer(0);
        setListenerTimer(0);
        setKeyTimer(0);
        setSignerTimer(0);
        setTimer(0);


        // @ts-ignore
        intervalBrokerRef.current = setInterval(() => {
            setBrokerTimer((brokerTimer) => brokerTimer + 1);
        }, 10);

        // @ts-ignore
        intervalListenerRef.current = setInterval(() => {
            setListenerTimer((listenerTimer) => listenerTimer + 1);
        }, 10);

        // @ts-ignore
        intervalKeyRef.current = setInterval(() => {
            setKeyTimer((keyTimer) => keyTimer + 1);
        }, 10);

        // @ts-ignore
        intervalSignerRef.current = setInterval(() => {
            setSignerTimer((signerTimer) => signerTimer + 1);
        }, 10);

        // @ts-ignore
        intervalRef.current = setInterval(() => {
            setTimer((timer) => timer + 1);
        }, 10);


        const {amountOfKeys, amountOfWorkers, batchSize, amountOfRecords} = values;
        const payload = {
            Action: "key",
            Key: {
                KeyLimit: amountOfKeys,
                BatchSize: batchSize,
                WorkersAmount: amountOfWorkers,
                RecordsAmount: amountOfRecords,
            }
        }
        setBrokerTotal(amountOfRecords + 1)
        setListenerTotal(amountOfRecords + 2)
        setKeyTotal(amountOfKeys)
        setSignerTotal(amountOfRecords)
        setLogsTotal(amountOfRecords)


        const response = await fetch('http://localhost:5003/handler', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(payload),
        });
    };

    const brokerMinutes = Math.floor((brokerTimer / 6000));
    const brokerSeconds = Math.floor((brokerTimer / 100) % 60);
    const brokerMilliseconds = (brokerTimer % 100).toString().padStart(2, '0');

    const listenerMinutes = Math.floor((listenerTimer / 6000));
    const listenerSeconds = Math.floor((listenerTimer / 100) % 60);
    const listenerMilliseconds = (listenerTimer % 100).toString().padStart(2, '0');

    const keyMinutes = Math.floor((keyTimer / 6000));
    const keySeconds = Math.floor((keyTimer / 100) % 60);
    const keyMilliseconds = (keyTimer % 100).toString().padStart(2, '0');

    const signerMinutes = Math.floor((signerTimer / 6000));
    const signerSeconds = Math.floor((signerTimer / 100) % 60);
    const signerMilliseconds = (signerTimer % 100).toString().padStart(2, '0');

    const minutes = Math.floor((timer / 6000));
    const seconds = Math.floor((timer / 100) % 60);
    const milliseconds = (timer % 100).toString().padStart(2, '0');


    useEffect(() => {
        if (!isSubmitting && logsTotal !== -1) {
            setBrokerTotal(-1)
            setListenerTotal(-1)
            setKeyTotal(-1)
            setSignerTotal(-1)
            setLogsTotal(-1)
        }
    }, [isSubmitting]);

    return (
        <ConfigProvider theme={{algorithm: theme.darkAlgorithm,}}>
            <Row justify="center" align="middle" style={{minHeight: '40vh'}}>
                <Col span={4}>
                    <Form form={form} onFinish={handleSubmit}>
                        <Form.Item label="Amount of Keys" name="amountOfKeys" initialValue={1}>
                            <InputNumber size="large" min={1} max={100}/>
                        </Form.Item>
                        <Form.Item label="Amount of Workers" name="amountOfWorkers" initialValue={1}>
                            <InputNumber size="large" min={1} max={100}/>
                        </Form.Item>
                        <Form.Item label="Batch Size" name="batchSize" initialValue={100}>
                            <InputNumber size="large" min={1} max={1000}/>
                        </Form.Item>
                        <Form.Item label="Amount of Records" name="amountOfRecords" initialValue={100}>
                            <InputNumber size="large" min={1} max={100000}/>
                        </Form.Item>
                        <Form.Item>
                            <Button size='large' type="primary" htmlType="submit"
                                    disabled={isSubmitting}>Submit</Button>
                        </Form.Item>
                    </Form>
                </Col>
            </Row>
            <div style={{width: '1400px', margin: '0 auto'}}>
                <Row gutter={16}>
                    <Col span={5}>
                        <Card>
                            <h1> {`${brokerMinutes.toString().padStart(2, '0')}:${brokerSeconds.toString().padStart(2, '0')}:${brokerMilliseconds}`}</h1>
                        </Card>
                    </Col>
                    <Col span={5}>
                        <Card>
                            <h1> {`${listenerMinutes.toString().padStart(2, '0')}:${listenerSeconds.toString().padStart(2, '0')}:${listenerMilliseconds}`}</h1>
                        </Card>
                    </Col>
                    <Col span={4}>
                        <Card>
                            <h1> {`${keyMinutes.toString().padStart(2, '0')}:${keySeconds.toString().padStart(2, '0')}:${keyMilliseconds}`}</h1>
                        </Card>
                    </Col>
                    <Col span={5}>
                        <Card>
                            <h1> {`${signerMinutes.toString().padStart(2, '0')}:${signerSeconds.toString().padStart(2, '0')}:${signerMilliseconds}`}</h1>
                        </Card>
                    </Col>
                    <Col span={5}>
                        <Card>
                            <h1> {`${minutes.toString().padStart(2, '0')}:${seconds.toString().padStart(2, '0')}:${milliseconds}`}</h1>
                        </Card>
                    </Col>

                </Row>
            </div>
            <Row justify="center" align="middle" style={{minHeight: '10vh'}}>
                <WebSocketListener
                    // Broker
                    setIsBrokerSubmitting={setIsBrokerSubmitting} intervalBrokerRef={intervalBrokerRef}
                    brokerTotal={brokerTotal}

                    // Listener
                    setIsListenerSubmitting={setIsListenerSubmitting} intervalListenerRef={intervalListenerRef}
                    listenerTotal={listenerTotal}

                    // Key
                    setIsKeySubmitting={setIsKeySubmitting} intervalKeyRef={intervalKeyRef} keyTotal={keyTotal}

                    // Signer
                    setIsSignerSubmitting={setIsSignerSubmitting} intervalSignerRef={intervalSignerRef}
                    signerTotal={signerTotal}

                    // Logger
                    setIsSubmitting={setIsSubmitting} intervalRef={intervalRef} logsTotal={logsTotal}
                />
            </Row>
            <Row justify="center" align="middle">
                <Row justify="center" align="middle" style={{width: '100%', height: '100px'}}>
                    <Switch onChange={toggleImage} checkedChildren="Animation" unCheckedChildren="Diagram"/>
                </Row>
                <Row justify="center" align="middle" style={{width: '100%'}}>
                    {isImageOne ? (
                        <img src="/signer.gif" style={{width: '1200px', height: '742px'}}/>
                    ) : (
                        <img src="/signer.animation.gif" style={{width: '1200px', height: '742px'}}/>
                    )}</Row>
            </Row>
        </ConfigProvider>
    )
}
