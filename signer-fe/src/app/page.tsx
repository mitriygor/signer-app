'use client';
import {Button, Col, ConfigProvider, Form, InputNumber, Row, theme} from 'antd';
// @ts-ignore
import WebSocketListener from "../../components/WebSocketListener";
// @ts-ignore
import {useState} from "react";


export default function Home() {
    const [form] = Form.useForm();


    const handleSubmit = async (values: any) => {
        console.log("handleSubmit")

        console.log("values", values)

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

        console.log("payload", payload)

        const response = await fetch('http://localhost:5003/handler', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(payload),
        });
        //
        // const data = await response.json();
        // console.log("data", data)
    };


    return (
        <ConfigProvider theme={{algorithm: theme.darkAlgorithm,}}>
            <Row justify="center" align="middle" style={{minHeight: '50vh'}}>
                <Col span={4}>
                    <Form form={form} onFinish={handleSubmit}>
                        <Form.Item label="Amount of Keys" name="amountOfKeys" initialValue={100}>
                            <InputNumber size="large" min={1} max={100}/>
                        </Form.Item>
                        <Form.Item label="Amount of Workers" name="amountOfWorkers" initialValue={100}>
                            <InputNumber size="large" min={1} max={100}/>
                        </Form.Item>
                        <Form.Item label="Batch Size" name="batchSize" initialValue={1000}>
                            <InputNumber size="large" min={1} max={1000}/>
                        </Form.Item>
                        <Form.Item label="Amount of Records" name="amountOfRecords" initialValue={100000}>
                            <InputNumber size="large" min={1} max={100000}/>
                        </Form.Item>
                        <Form.Item>
                            <Button size='large' type="primary" htmlType="submit">Submit</Button>
                        </Form.Item>
                    </Form>
                </Col>
            </Row>
            <WebSocketListener/>
        </ConfigProvider>
    )
}
