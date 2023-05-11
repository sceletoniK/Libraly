import React, {useState, useEffect} from 'react'
import {List, Button, Input, Space, Col, Row, Typography, Divider, Form, Card, theme, Tag} from 'antd';
import {Link} from "react-router-dom";
import axios from "axios";
import {CloseOutlined, FileTextOutlined} from "@ant-design/icons";

const {Title, Text} = Typography;
const {CheckableTag} = Tag;


const Books: React.FC = () => {

    const [error, setError] = useState<Error | null>(null);
    const [isLoaded, setIsLoaded] = useState(false);
    const [items, setItems] = useState<BookType[]>([]);

    const [selectedTags, setSelectedTags] = useState<string[]>([]);
    const [genres, setGenres] = useState<GenreType[]>([]);

    interface BookType {
        id: React.Key,
        name: string,
        author: string,
        publisher: string,
    }


    useEffect(() => {
        fetch("http://localhost:8080/book")
            .then(res => res.json())
            .then(
                (result) => {
                    setIsLoaded(true);
                    setItems(result);
                },
                (error) => {
                    setIsLoaded(true);
                    setError(error);
                }
            );
        axios.get(
            "http://localhost:8080/genre"
        ).then((result) => {
            setGenres(result.data)
        }).catch((error) => {
            console.log('Пятисотая на проде!')
        })
    }, [])


    interface GenreType {
        genreId: React.Key,
        name: string,
    }


    const handleChange = (tag: string, checked: boolean) => {
        const nextSelectedTags = checked
            ? [...selectedTags, tag]
            : selectedTags.filter((t) => t !== tag);
        setSelectedTags(nextSelectedTags);
    };

    if (error) {
        return <div>Ошибка: {error.message}</div>;
    } else if (!isLoaded) {
        return <div>Загрузка...</div>;
    } else {
        return (
            <Row justify='center' gutter={40} style={{height: '100%'}}>
                <Col span={4} >
                    <Space align={'start'}>
                    <Form
                        name="basic"
                        labelCol={{span: 8}}
                        wrapperCol={{span: 18}}
                        initialValues={{remember: true}}
                        autoComplete="off"
                    >
                        <Title level={3}>Фильтр</Title>
                        <Form.Item
                            label="Название"
                            name="name"
                        >
                            <Input/>
                        </Form.Item>

                        <Form.Item
                            label="Автор"
                            name="author"
                        >
                            <Input/>
                        </Form.Item>

                        <Form.Item
                            label="Издатель"
                            name="pubisher"
                        >
                            <Input/>
                        </Form.Item>
                        <Space size={[0, 8]} wrap>
                            {genres.map((tag) => (
                                <CheckableTag
                                    key={tag.genreId}
                                    checked={selectedTags.includes(tag.name)}
                                    onChange={(checked) => handleChange(tag.name, checked)}
                                >
                                    {tag.name}
                                </CheckableTag>
                            ))}
                        </Space>
                        <br/>
                        <br/>
                        <Form.Item wrapperCol={{offset: 4, span: 16}}>
                            <Space size='large'>
                                <Button type="primary" htmlType="submit">
                                    Поиск
                                </Button>
                                <Button htmlType="button">
                                    Очистить
                                </Button>
                            </Space>
                        </Form.Item>
                    </Form>
                    </Space>
                </Col>
                <Col span={20} >
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
                        dataSource={items}
                        renderItem={(item: BookType) => (
                            <List.Item
                                key={item.id}
                            >
                                <Link to={`/book/${item.id}`}>
                                    <Card
                                        hoverable
                                        style={{width: 300}}
                                        cover={<img
                                            width={272}
                                            alt="logo"
                                            src="https://gw.alipayobjects.com/zos/rmsportal/mqaQswcyDLcXyDKnZfES.png"
                                        />}
                                    >
                                        <Title level={2}>{item.name}</Title>
                                        <Space direction="vertical">
                                            <Text>{item.author}</Text>
                                            <Text type="secondary">{item.publisher}</Text>
                                        </Space>
                                    </Card>
                                </Link>
                            </List.Item>
                        )}/>
                </Col>
            </Row>
        );
    }
}

export default Books;