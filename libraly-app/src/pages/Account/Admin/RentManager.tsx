import React from "react";
import moment from "moment/moment";
import {Button, Popconfirm, Table, Tag} from "antd";
import {Link} from "react-router-dom";

interface RentType {
    id: React.Key,
    book: string,
    instanceId: number
    client: string
    status: React.ReactNode
    requestDate: moment.Moment
    startRentDate: moment.Moment
    deadline: moment.Moment
}

const originData: RentType[] = [
    {
        id: 1,
        book: 'Библия программиста',
        instanceId: 7,
        client: 'admin',
        status: <Tag color="success">Одобрено</Tag>,
        requestDate: moment(),
        startRentDate: moment(),
        deadline: moment()

    },
    {
        id: 2,
        book: 'Наедине с собой',
        instanceId: 34,
        client: 'admin',
        status: <Tag color="processing">На рассмотрении</Tag>,
        requestDate: moment(),
        startRentDate: moment(new Date(2000, 0, 1, 3, 0, 0)),
        deadline: moment(new Date(2000, 0, 1, 3, 0, 0))
    },
];


const RentManager = () => {

    const columns = [
        {
            title: 'Название',
            dataIndex: 'book',
            width: '19%',
            render: (_: any, record: RentType) => {return <Link to={'/book/38'}>{record.book}</Link>}
        },
        {
            title: 'Клиент',
            dataIndex: 'client',
            width: '19%',
        },
        {
            title: 'Код книги',
            dataIndex: 'instanceId',
            width: '19%',
        },
        {
            title: 'Статус',
            dataIndex: 'status',
            width: '19%',
        },
        {
            title: 'Операции',
            dataIndex: 'operation',
            render: (_: any, record: RentType) => {
                if (record.startRentDate.isSame(moment(new Date(2000, 0, 1, 3, 0, 0))))
                    return (
                        <span>
                            <Popconfirm title="Одобрить аренду ?" >
                              <Button type="primary" size={'small'} style={{marginRight: 8}}>
                                Одобрить
                              </Button>
                            </Popconfirm>
                            <Popconfirm title="Отказать в аренде ?">
                              <Button type="primary" size={'small'}>
                                Отказать
                              </Button>
                            </Popconfirm>
                        </span>
                    )
                else
                    return (
                        <span>
                            <Popconfirm title="Закрыть аренду ?">
                              <Button type="primary" size="small">
                                  Закрыть
                              </Button>
                            </Popconfirm>
                        </span>
                    )
            },
        },
    ];

    return <Table bordered columns={columns} dataSource={originData} />
}

export default RentManager