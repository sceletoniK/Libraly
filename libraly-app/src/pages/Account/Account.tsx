import React, {useEffect, useState, useMemo} from 'react';
import {useNavigate, NavigateOptions} from "react-router-dom";
import axios from "axios";
import {Col, Row, Avatar, Typography, Tabs, Space, Button} from 'antd';
import {AxiosError} from "axios";
import {UserOutlined, AndroidOutlined, AppleOutlined} from '@ant-design/icons';
import './Account.css';
import type { TabsProps } from 'antd';
import Cart from "./Cart";
import Rent from "./Rent";
import History from "./History";
import BookManager from "./Admin/BookManager";
import RentManager from "./Admin/RentManager";

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
                    navigate('/refresh')
                } else {
                    console.log(error);
                    setError(error);
                }
            })
    }, [navigate]);

    const tabs: TabsProps['items'] = window.location.pathname === '/admin' ?
        [
            {
                key: '4',
                label: `Менеджер книг`,
                children: <BookManager />,
            },
            {
                key: '5',
                label: `Менеджер аренд`,
                children: <RentManager />,
            },
            {
                key: '6',
                label: `Менеджер жанров`,
                children: <></>,
            },
            {
                key: '7',
                label: `Менеджер пользователей`,
                children: <></>,
            },

        ] : [
            {
                key: '1',
                label: `Корзина`,
                children: <Cart/>,
            },
            {
                key: '2',
                label: `Аренды`,
                children: <Rent/>,
            },
            {
                key: '3',
                label: `История`,
                children: <History />,
            },
    ];

    if (error) {
        return <div>Технические шоколадки, зайдите попозже :З</div>;
    } else if (!isLoaded) {
        return <div>Загрузка...</div>;
    } else {
        return (
            <>
                <Row justify="space-around" align="top">
                    <Col sm={{span: 24}} lg={{span: 6}} className='avatar_column'>
                        <Avatar size={192} icon={<UserOutlined/>}/>
                        <Title level={2}>{items['login']}</Title>
                        <Text code type="success" style={{fontSize: 18}}>Задолженностей нет</Text>
                        <br/>
                        {items['admin'] && (
                            <Text>Администратор</Text>
                        )}
                        {items['admin'] && window.location.pathname === '/account' && (
                            <>
                                <br/>
                                <br/>
                                <Button size={'large'} type={'primary'} onClick={() => navigate('/admin')}>
                                    Панель администратора
                                </Button>
                            </>
                        )}
                        {items['admin'] && window.location.pathname === '/admin' && (
                            <>
                                <br/>
                                <br/>
                                <Button size={'large'} type={'primary'} onClick={() => navigate('/account')}>
                                    Панель пользователя
                                </Button>
                            </>
                        )}
                        <br/>
                        <br/>
                    </Col>
                    <Col sm={{span: 24}} lg={{span: 16}} className='tabs_column'>
                        <Tabs
                            defaultActiveKey="1"
                            centered
                            size="large"
                            type="card"
                            className='tabs'
                            items={tabs}
                        />
                    </Col>
                </Row>
            </>
        )
    }
}

export default Account;