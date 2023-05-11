import React, {useEffect, useState} from 'react'
import {Link, useNavigate} from "react-router-dom";
import axios, {AxiosError} from "axios";
import moment from "moment";
import {Row, Col, Table} from 'antd'

interface HistoryType{
    ClientId: number
    BookId: number
    StartRentDate: string
    EndRentDate: string
}

const History = () => {
    const navigate = useNavigate();
    const [isLoaded, setIsLoaded] = useState(false);
    const [items, setItems] = useState<HistoryType[]>([{
        BookId: 38,
        ClientId: 8,
        StartRentDate: moment().format('DD/MM/YYYY'),
        EndRentDate: moment().format('DD/MM/YYYY')
    }
    ]);
    const [error, setError] = useState<AxiosError | null>(null);


    const columns = [
        {
            title: 'Книга',
            dataIndex: 'BookId',
            key: 'BookId',
            render:  () => <Link to={`/book/38`}>Наедине с собой</Link>
        },
        {
            title: 'Начало аренды',
            dataIndex: 'StartRentDate',
            key: 'startRentDate',
        },
        {
            title: 'Конец аренды',
            dataIndex: 'EndRentDate',
            key: 'EndRentDate',
        },
    ];

    return(
        <Row justify={'center'} align={'middle'}>
            <Col span={22}>
                <Table columns={columns} dataSource={items} />
            </Col>
        </Row>
    )
}

export default History