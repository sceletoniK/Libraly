import React, {useRef, useState} from "react";
import {Row, Col, Space, Card, Typography, Table, Button, Popconfirm, InputRef, Input, Divider} from 'antd'
import moment from "moment/moment";
import {ColumnsType, ColumnType} from "antd/es/table";
import {FilterConfirmProps} from "antd/es/table/interface";
import {SearchOutlined} from "@ant-design/icons";

const {Title, Text} = Typography

interface BookType {
    id: React.Key,
    name: string,
    author: string,
    publisher: string,
    publishDate: moment.Moment
}
type DataIndex = keyof BookType;

const BookData: BookType[] = [
    {
        id: 1,
        name: 'Наедине с собой',
        author: 'Марк Аврелий',
        publisher: 'АСТ',
        publishDate: moment()
    }
];

interface InstanceType{
    id: React.Key
    instanceId: number
}

const InstanceData: InstanceType[] = [
    {
        id: 1,
        instanceId: 46,
    },
    {
        id: 2,
        instanceId: 11,
    },
    {
        id: 3,
        instanceId: 67,
    },
    {
        id: 4,
        instanceId: 123,
    },
    {
        id: 5,
        instanceId: 45,
    },
]

const InstanceManager = () => {
    const [searchText, setSearchText] = useState('');
    const [searchedColumn, setSearchedColumn] = useState('');
    const searchInput = useRef<InputRef>(null);

    const handleSearch = (
        selectedKeys: string[],
        confirm: (param?: FilterConfirmProps) => void,
        dataIndex: DataIndex,
    ) => {
        confirm();
        setSearchText(selectedKeys[0]);
        setSearchedColumn(dataIndex);
    };

    const handleReset = (clearFilters: () => void) => {
        clearFilters();
        setSearchText('');
    };

    const getColumnSearchProps = (dataIndex: DataIndex): ColumnType<BookType> => ({
        filterDropdown: ({ setSelectedKeys, selectedKeys, confirm, clearFilters, close }) => (
            <div style={{ padding: 8 }} onKeyDown={(e) => e.stopPropagation()}>
                <Input
                    ref={searchInput}
                    placeholder={`Search ${dataIndex}`}
                    value={selectedKeys[0]}
                    onChange={(e) => setSelectedKeys(e.target.value ? [e.target.value] : [])}
                    onPressEnter={() => handleSearch(selectedKeys as string[], confirm, dataIndex)}
                    style={{ marginBottom: 8, display: 'block' }}
                />
                <Space>
                    <Button
                        type="primary"
                        onClick={() => handleSearch(selectedKeys as string[], confirm, dataIndex)}
                        icon={<SearchOutlined />}
                        size="small"
                        style={{ width: 90 }}
                    >
                        Search
                    </Button>
                    <Button
                        onClick={() => clearFilters && handleReset(clearFilters)}
                        size="small"
                        style={{ width: 90 }}
                    >
                        Reset
                    </Button>
                    <Button
                        type="link"
                        size="small"
                        onClick={() => {
                            close();
                        }}
                    >
                        close
                    </Button>
                </Space>
            </div>
        ),
        filterIcon: (filtered: boolean) => (
            <SearchOutlined style={{ color: filtered ? '#1890ff' : undefined }} />
        ),
        onFilter: (value, record) =>
            record[dataIndex]
                .toString()
                .toLowerCase()
                .includes((value as string).toLowerCase()),
        onFilterDropdownOpenChange: (visible) => {
            if (visible) {
                setTimeout(() => searchInput.current?.select(), 100);
            }
        },
    });

    const BookColumns: ColumnsType<BookType> = [
        {
            title: 'Название',
            dataIndex: 'name',
            width: '20%',
            filterMode: 'tree',
            filterSearch: true,
            ...getColumnSearchProps('name'),
        },
        {
            title: 'Автор',
            dataIndex: 'author',
            width: '20%',
            ...getColumnSearchProps('author'),
        },
        {
            title: 'Издатель',
            dataIndex: 'publisher',
            ...getColumnSearchProps('publisher'),
            width: '20%',
        },
        {
            title: 'Дата издания',
            dataIndex: 'publishDate',
            width: '20%',
            ...getColumnSearchProps('publishDate'),
            render: (value: moment.Moment) => value.format('DD/MM/YYYY')
        },
        {
            title: '',
            dataIndex: 'operation',
            render: (_: any, record: BookType) => {
                return (
                    <span>
                        <Typography.Link>
                            <Button type="primary" size="small">
                                Копии
                            </Button>
                        </Typography.Link>
                    </span>
                )
            },
        },
    ];

    const InstanceColumns: ColumnsType<InstanceType> = [
        {
            title: 'Код книги',
            dataIndex: 'instanceId',
            filterMode: 'tree',
            filterSearch: true,
            onFilter: (value, record) => record.instanceId === value
        },
        {
            title: '',
            dataIndex: 'operation',
            render: (_: any, record: InstanceType) => {
                return (
                    <Popconfirm title="Удалить данную копию ?">
                        <Button type="primary" size="small">
                            Удалить
                        </Button>
                    </Popconfirm>
                )
            }
        }
    ]

    return (
        <>
            <Row justify={'center'} gutter={[0, 40]}>
                <Col span={8}>
                    <Card
                        hoverable
                        style={{width: 300}}
                        cover={<img
                            width={272}
                            alt="logo"
                            src="https://gw.alipayobjects.com/zos/rmsportal/mqaQswcyDLcXyDKnZfES.png"
                        />}
                    >
                        <Title level={2}>Смерть Артура</Title>
                        <Space direction="vertical">
                            <Text>Томас Мэлори</Text>
                            <Text type="secondary">Большие Книги</Text>
                        </Space>
                    </Card>
                </Col>
                <Col span={6}>
                        <Table
                            bordered
                            dataSource={InstanceData}
                            columns={InstanceColumns}
                            pagination={{ pageSize: 6 }}
                            size={'small'}
                        />
                        <Button type={'primary'}>
                            Добавить новую копию
                        </Button>
                </Col>
            </Row>
            <Divider/>
            <Row>
                <Col span={24}>
                    <Table
                        bordered
                        dataSource={BookData}
                        columns={BookColumns}
                    />
                </Col>
            </Row>
        </>
    )
}

export default InstanceManager