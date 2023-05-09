import React, {useEffect, useState} from 'react';
import axios from "axios";
import {Link, useNavigate} from "react-router-dom";
import {Button, Card, List, Space, Typography, Col, Row} from "antd";
import {
    CloseOutlined, FileTextOutlined, ShoppingCartOutlined
} from '@ant-design/icons';

const {Text, Title} = Typography

interface BookType {
    id: React.Key,
    name: string,
    author: string,
    publisher: string,
}


const Cart = () => {
    const navigate = useNavigate();
    const [cart, setCart] = useState<BookType[]>([]);

    useEffect(() => {
        axios.get(
            'http://localhost:8080/cart',
            {
                headers: {
                    Authorization: `Bearer ${localStorage.getItem('token')}`
                }
            }
        ).then((result) => {
            if (result.data.length)
                setCart(result.data)
            else
                setCart([])
        }).catch((error) => {
            if (error.response && error.response.status === 401) {
                navigate('/login');
            } else if (error.response && error.response.status === 901) {
                navigate('/refresh')
            } else {
                console.log(error);
            }
        })
    }, [navigate])

    return (
        <>
        <Row justify={'center'} align={'middle'} gutter={[0, 8]}>
            <Col span={22}>
            <List
                itemLayout="vertical"
                size="large"
                pagination={{
                    pageSize: 12,
                    position: 'bottom',
                    align: 'center'
                }}
                grid={{
                    gutter: 20,
                    xs: 1,
                    sm: 2,
                    md: 2,
                    lg: 3,
                    xl: 3,
                    xxl: 3,
                }}
                dataSource={cart}
                renderItem={(item: BookType) => (
                    <List.Item
                        key={item.id}
                    >
                        <Card
                            hoverable
                            style={{width: 300}}
                            cover={<img
                                width={272}
                                alt="logo"
                                src="https://gw.alipayobjects.com/zos/rmsportal/mqaQswcyDLcXyDKnZfES.png"
                            />}
                            actions={[(
                                <Button type={'text'}>
                                    <Link to={`/book/${item.id}`}>
                                        <FileTextOutlined/> Подробнее
                                    </Link>
                                </Button>),
                                (<Button type={'text'}><CloseOutlined/> Убрать</Button>)
                            ]}
                        >
                            <Title level={2}>{item.name}</Title>
                            <Space direction="vertical">
                                <Text>{item.author}</Text>
                                <Text type="secondary">{item.publisher}</Text>
                            </Space>
                        </Card>
                    </List.Item>
                )}/>
            </Col>
        </Row>
            <br/>
        <Row justify={'center'} align={'middle'} gutter={[0, 8]}>
            <Col span={6} style={{display: 'flex', justifyContent:'center'}}>
                <Button type="primary" icon={<ShoppingCartOutlined />} size={'large'}>
                    Оформить
                </Button>
            </Col>
        </Row>
        </>
    )
}

export default Cart