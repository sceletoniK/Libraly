import React, {useState, useRef} from "react";
import {Form, Input, InputNumber, Popconfirm, Table,
    Typography, DatePicker, Button, Divider, Upload, Row, Col, Space} from 'antd';
import type { InputRef } from 'antd';
import type { ColumnsType, ColumnType } from 'antd/es/table';
import {InboxOutlined, SearchOutlined} from '@ant-design/icons'
import type { FilterConfirmProps } from 'antd/es/table/interface';
import moment from "moment";

const { TextArea } = Input;
interface BookType {
    id: React.Key,
    name: string,
    author: string,
    publisher: string,
    publishDate: moment.Moment
}

type DataIndex = keyof BookType;

const originData: BookType[] = [
    {
        id: 1,
        name: 'Наедине с собой',
        author: 'Марк Аврелий',
        publisher: 'АСТ',
        publishDate: moment()
    }
];

interface EditableCellProps extends React.HTMLAttributes<HTMLElement> {
    adding: boolean
    editing: boolean;
    dataIndex: string;
    title: any;
    inputType: 'number' | 'text' | 'date';
    record: BookType;
    index: number;
    children: React.ReactNode;
}

const EditableCell: React.FC<EditableCellProps> = ({
                                                       adding,
                                                       editing,
                                                       dataIndex,
                                                       title,
                                                       inputType,
                                                       record,
                                                       index,
                                                       children,
                                                       ...restProps
                                                   }) => {
    const inputNode = inputType === 'number' ? <InputNumber/> : inputType === 'date' ? <DatePicker/> : <Input/>;

    return (
        <td {...restProps}>
            {editing || adding ? (
                <Form.Item
                    name={dataIndex}
                    style={{margin: 0}}
                    rules={[
                        {
                            required: true,
                            message: `Укажите поле: ${title}!`,
                        },
                    ]}
                >
                    {inputNode}
                </Form.Item>
            ) : (
                children
            )}
        </td>
    );
};

const BookManager = () => {

    const [form] = Form.useForm();
    const [data, setData] = useState(originData);
    const [editingKey, setEditingKey] = useState<number | string>(0);

    const isEditing = (record: BookType) => record.id === editingKey;

    const edit = (record: Partial<BookType> & { id: React.Key }) => {
        form.setFieldsValue({name: '', author: '', publisher: '', publishDate: '', ...record});
        setEditingKey(record.id);
    };

    const cancel = () => {
        setEditingKey(0);
    };

    const save = async (key: React.Key) => {
        try {
            const row = (await form.validateFields()) as BookType;

            const newData = [...data];
            const index = newData.findIndex((item) => key === item.id);
            if (index > 0) {
                console.log(index)
                const item = newData[index];
                newData.splice(index, 1, {
                    ...item,
                    ...row,
                });
                setData(newData);
                setEditingKey(0);
            } else {
                newData.push(row);
                setData(newData);
                setEditingKey(0);
            }
        } catch (errInfo) {
            console.log('Validate Failed:', errInfo);
        }
    };

    const handleDelete = (key: React.Key) => {
        const newData = data.filter((item) => item.id !== key);
        setData(newData);
    };

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

    const columns: ColumnsType<BookType> = [
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
                const editable = isEditing(record);
                if (editable)
                    return (
                        <span>
                            <Typography.Link onClick={() => save(record.id)} style={{marginRight: 8}}>
                              <Button type="primary" size="small">
                                Сохранить
                              </Button>
                            </Typography.Link>
                            <Popconfirm title="Sure to cancel?" onConfirm={cancel}>
                              <Button type="primary" size="small">
                                Отмена
                              </Button>
                            </Popconfirm>
                        </span>
                    )
                else if (record.id === -1)
                    return (
                        <span>
                            <Popconfirm title="Sure to add?" onConfirm={() => save(record.id)}>
                              <Button type="primary" size="small">
                                Добавить
                              </Button>
                            </Popconfirm>
                        </span>
                    )
                else
                    return (
                        <span>
                            <Typography.Link disabled={editingKey !== 0} onClick={() => edit(record)} style={{marginRight: 8}}>
                                <Button type="primary" size="small">
                                    Изменить
                                </Button>
                            </Typography.Link>
                            <Popconfirm title="Sure to delete?" onConfirm={() => handleDelete(record.id)}>
                                <Button type="primary" size="small">
                                    Удалить
                              </Button>
                            </Popconfirm>
                        </span>
                    )
            },
        },
    ];

    return (
        <>
            <Form
                labelCol={{ span: 8 }}
                wrapperCol={{ span: 16 }}
            >
                <Row justify={'center'} gutter={40}>
                    <Col span={8} style={{maxWidth: 300}}>
                        <Form.Item name="dragger" valuePropName="fileList" noStyle>
                            <Upload.Dragger name="files" action="/upload.do">
                                <p className="ant-upload-drag-icon">
                                    <InboxOutlined />
                                </p>
                                <p className="ant-upload-text">Перетащите или кликните чтобы загрузить обложку</p>
                                <p className="ant-upload-hint">Обложка должна быть одна и в формате JPG\PNG</p>
                            </Upload.Dragger>
                        </Form.Item>
                    </Col>
                    <Col span={8}>
                        <br/>
                        <Form.Item
                            name="Название"
                            label="Название"
                            rules={[{ required: true}]}
                            style={{maxWidth: 350}}
                        >
                            <Input placeholder="Введние название книги" />
                        </Form.Item>
                        <Form.Item
                            name="Автор"
                            label="Автор"
                            rules={[{ required: true}]}
                            style={{maxWidth: 350}}
                        >
                            <Input placeholder="Введние автора книги" />
                        </Form.Item>
                        <Form.Item
                            name="Издатель"
                            label="Издатель"
                            rules={[{ required: true}]}
                            style={{maxWidth: 350}}
                        >
                            <Input placeholder="Введние издателя книги" />
                        </Form.Item>
                        <Form.Item
                            name="Дата издания"
                            label="Дата издания"
                            style={{maxWidth: 400}}
                        >
                            <DatePicker placeholder="Выберите дату издания" style={{width: '100%'}} />
                        </Form.Item>
                    </Col>
                    <Col span={8}>
                        <br/>
                        <Form.Item label="Описание" required>
                            <TextArea rows={12} autoSize={{ minRows: 10, maxRows: 10 }} placeholder='Введите описание книги'/>
                        </Form.Item>
                    </Col>
                </Row>
                <Row justify={'center'}>
                    <Col span={6}>
                        <Button type={'primary'} size={'large'}>
                            Добавить книгу
                        </Button>
                    </Col>
                </Row>
            </Form>
            <Divider/>
                <Table
                    bordered
                    dataSource={data}
                    columns={columns}
                    pagination={{
                        onChange: cancel,
                    }}

                />
        </>
    )
}

export default BookManager