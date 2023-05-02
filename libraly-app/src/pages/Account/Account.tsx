import React, {useEffect, useState} from 'react';
import {useNavigate} from "react-router-dom";
import axios from "axios";
import {Col, Row, Avatar, Typography, Tabs,} from 'antd';
import {AxiosError} from "axios";
import {UserOutlined, AndroidOutlined, AppleOutlined} from '@ant-design/icons';
import './Account.css';

const {Title, Text} = Typography;

const Account: React.FC = () => {

    const navigate = useNavigate();
    const [isLoaded, setIsLoaded] = useState(false);
    const [items, setItems] = useState<any>([]);
    const [error, setError] = useState<AxiosError | null>(null);

    useEffect(() => {
        axios.get("http://localhost:8080/user",
            {
                headers: {
                    Authorization: `Bearer ${localStorage.getItem('token')}`
                }
            })
            .then((result) => {
                setItems(result.data);
                setIsLoaded(true);
            })
            .catch((error: AxiosError) => {
                if (error.response && error.response.status === 401) {
                    navigate('/login');
                } else if (error.response && error.response.status === 901) {
                    console.log('Просрочено');
                    navigate('/reg');
                } else {
                    console.log(error);
                    setError(error);
                }
            })
    }, []);

    if (error) {
        return <div>Технические шоколадки, зайдите попозже :З</div>;
    } else if (!isLoaded) {
        return <div>Загрузка...</div>;
    } else {
        return (
            <>
                <Row justify="space-around" align="top">
                    <Col span={4} className='avatar_column'>
                        <Avatar size={192} icon={<UserOutlined/>}/>
                        <Title level={2}>{items['login']}</Title>
                        {items['admin'] && (
                            <Text>Администратор</Text>
                        )}
                    </Col>
                    <Col span={16} className='tabs_column'>
                        <Tabs
                            defaultActiveKey="1"
                            centered
                            size="large"
                            type="card"
                            className='tabs'
                            items={[AppleOutlined, AndroidOutlined].map((Icon, i) => {
                                const id = String(i + 1);

                                return {
                                    label: (
                                        <span>
                                            <Icon/>
                                            Tab {id}
                                        </span>
                                    ),
                                    key: id,
                                    children: `Tab ${id}`,
                                };
                            })}
                        />
                    </Col>
                </Row>
            </>
        )
    }
}

export default Account;