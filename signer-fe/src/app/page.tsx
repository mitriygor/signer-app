'use client';
import {Button, Col, ConfigProvider, Form, InputNumber, Row, theme} from 'antd';
// @ts-ignore
import WebSocketListener from "../../components/WebSocketListener";
// @ts-ignore
import {useState} from "react";


export default function Home() {
    const [formState, setFormState] = useState({
        amountOfKeys: 100,
        amountOfWorkers: 100,
        batchSize: 1000,
        amountOfRecords: 100000,
    });


    const handleSubmit = async (e: any) => {
        e.preventDefault();

        const response = await fetch('/api/your-endpoint', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(formState),
        });

        const data = await response.json();
    };


    return (
        <ConfigProvider theme={{algorithm: theme.darkAlgorithm,}}>
            <Row justify="center" align="middle" style={{minHeight: '50vh'}}>
                <Col span={4}>
                    <Form>
                        <Form.Item label="Amount of Keys" name="layout">
                            <InputNumber size="large" min={1} max={100} value={formState.amountOfKeys}/>
                        </Form.Item>
                        <Form.Item label="Amount of Workers" name="layout">
                            <InputNumber size="large" min={1} max={100} value={formState.amountOfWorkers}/>
                        </Form.Item>
                        <Form.Item label="Batch Size" name="layout">
                            <InputNumber size="large" min={1} max={1000}/>
                        </Form.Item>
                        <Form.Item label="Amount of Records" name="layout">
                            <InputNumber size="large" min={1} max={100000}/>
                        </Form.Item>
                        <Button size='large' type="primary">Send</Button>
                    </Form>
                </Col>
            </Row>
            <WebSocketListener/>
        </ConfigProvider>
    )
}
