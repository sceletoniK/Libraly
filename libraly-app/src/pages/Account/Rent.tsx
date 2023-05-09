import React, {useEffect, useState} from 'react';
import axios from "axios";
import {Link, useNavigate} from "react-router-dom";
import {Button, Card, List, Space, Typography, Col, Row, Table, Tag} from "antd";
import {
    CloseOutlined, FileTextOutlined, ShoppingCartOutlined
} from '@ant-design/icons';
import moment from "moment";
import {AxiosError} from "axios/index";
import {stat} from "fs";
import {ALL} from "dns";
const {Text, Title} = Typography
const { Column, ColumnGroup } = Table;

interface RentType {
    key: number
    name: string
    Status: React.ReactNode
    RequestDate: string
    StartRentDate: string
    deadline: string
}

interface RawRentType{
    instanceId: number
    clientId: number
    RequestDate: string
    startRentDate: string
    deadline: string
}

const Rent = () => {

    const navigate = useNavigate();
    const [isLoaded, setIsLoaded] = useState(false);
    const [items, setItems] = useState<RentType[]>([]);
    const [error, setError] = useState<AxiosError | null>(null);

    useEffect(() => {
        axios.get("http://localhost:8080/rent",
            {
                headers: {
                    Authorization: `Bearer ${localStorage.getItem('token')}`
                }
            })
            .then((result) => {
                const rawData: RawRentType[] = result.data
                let data: RentType[] = []
                rawData.map((rawVal, index) => {
                    const requestDate = moment(rawVal.RequestDate);
                    const startRentDate = moment(rawVal.startRentDate);
                    const deadline = moment(rawVal.deadline);
                    let status = <></>
                    if(startRentDate.isSame(moment(new Date(2000, 0, 1, 3, 0, 0))))
                        status = <Tag color="processing">На рассмотрении</Tag>
                    else if(!deadline.isSame(moment(new Date(2000, 0, 1, 3, 0, 0))) && moment() > deadline)
                        status = <Tag color="error">Просрочено</Tag>
                    else
                        status = <Tag color="success">Одобрено</Tag>
                    data.push({
                        key: index,
                        Status: status,
                        deadline: !moment(rawVal.startRentDate).isSame(moment(new Date(2000, 0, 1, 3, 0, 0)))
                            ? moment(rawVal.deadline).format('DD/MM/YYYY') : '-',
                        name: rawVal.instanceId.toString(),
                        RequestDate: requestDate.format('DD/MM/YYYY'),
                        StartRentDate: !moment(rawVal.startRentDate).isSame(moment(new Date(2000, 0, 1, 3, 0, 0)))
                            ? moment(rawVal.startRentDate).format('DD/MM/YYYY') : '-'
                    })
                })
                console.log(data)
                setItems(data);
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

    const columns = [
        {
            title: 'Книга',
            dataIndex: 'name',
            key: 'name',
            render:  () => <Link to={`/books`}>Наедине с собой</Link>
        },
        {
            title: 'Статус',
            dataIndex: 'Status',
            key: 'status',
        },
        {
            title: 'Дата запроса',
            dataIndex: 'RequestDate',
            key: 'requestDate',
        },
        {
            title: 'Начало аренды',
            dataIndex: 'StartRentDate',
            key: 'startRentDate',
        },
        {
            title: 'Крайний срок',
            dataIndex: 'deadline',
            key: 'deadline',
        }
    ];

    if (error) {
        return <div>Ошибка: {error.message}</div>;
    } else if (!isLoaded) {
        return <div>Загрузка...</div>;
    } else {
        return(
            <Row justify={'center'} align={'middle'} gutter={[0, 8]}>
                <Col span={22}>
                    <Table columns={columns} dataSource={items}/>
                </Col>
            </Row>
        )
    }

}

export default Rent