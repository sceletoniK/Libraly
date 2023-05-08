import React, {useEffect, useState, useMemo} from 'react';
import {useSearchParams} from "react-router-dom";
import axios, {AxiosError} from "axios";
import {useParams} from "react-router-dom";
import {Row, Col, Image, Space, Typography, Divider, Tag, Button, Descriptions} from 'antd'

const {Title, Text, Paragraph} = Typography

const Book = () => {
    let { id } = useParams();


    return(
        <Row justify={'center'}>
            <Col md={{ span: 24, order: 2 }} xl={{ span: 12, order: 2 }} xxl={{span: 8, order: 2}}>
                <Space direction={'vertical'} align={'center'} size={'large'}>
                    <Image
                        width={450}
                        src="https://gw.alipayobjects.com/zos/rmsportal/mqaQswcyDLcXyDKnZfES.png"
                    />
                    <Text code type="success" style={{fontSize: 24}}>В наличии</Text>
                    <Button type="primary" shape="round"  size={'large'}>
                        Добавить в корзину
                    </Button>
                </Space>
            </Col>
            <Col md={{ span: 24, order: 1 }} xl={{ span: 12, order: 2 }} xxl={{span: 16, order: 2}}>
                    <Title>Библия программиста</Title>
                    <Space size={[0, 8]} wrap>
                        <Tag color="red">Классика</Tag>
                        <Tag color="gold">Фентези</Tag>
                        <Tag color="green">Детское</Tag>
                        <Tag color="blue">Зарубежное</Tag>
                    </Space>
                    <br/><br/>
                    <Descriptions column={1}>
                        <Descriptions.Item label="Автор">Аноним</Descriptions.Item>
                        <Descriptions.Item label="Издатель">Пикабу</Descriptions.Item>
                    </Descriptions>
                    <Divider plain style={{margin: 3}}/>
                    <Title level={3}>Описание</Title>

                        <Paragraph>
                            00. В начале было слово, и слово было 2 байта, а больше ничего не было.
                        </Paragraph>
                        <Paragraph>
                            01. И отделил Бог единицу от нуля, и увидел, что это хорошо.
                        </Paragraph>
                        <Paragraph>
                            02. И сказал Бог: да будут данные, и стало так.
                        </Paragraph>
                        <Paragraph>
                            03. И сказал Бог: да соберутся данные каждые в свое место, и создал дискеты, и винчестеры, и компакт-диски.
                        </Paragraph>
                        <Paragraph>
                            04. И сказал Бог: да будут компьютеры, чтобы было куда пихать дискеты, и винчестеры, и компакт-диски, и сотворил компьютеры, и нарек их хардом, и отделил хард от софта.
                        </Paragraph>
                        <Paragraph>
                            05. Софта же еще не было, но Бог быстро исправился, и создал программы большие и маленькие, и сказал им: плодитесь и размножайте, и заполняйте всю память.
                        </Paragraph>
                        <Paragraph>
                            06. Но надоело Ему создавать программы самому, и сказал Бог: создадим программиста по образу и подобию нашему, и да владычествует над компьютерами, и над программами, и над данными. И создал Бог программиста, и поселил его в своем ВЦ, чтобы работал в нем. И повел Он программиста к дереву каталогов, и заповедал: из всякого каталога можешь запускать программы, только из каталога Windоws не запускай, ибо маст дай (must die).
                        </Paragraph>
                        <Paragraph>
                            07. И сказал Бог: не хорошо программисту быть одному, сотворим ему пользователя, соответственно ему. И взял Он у программиста кость, в кой не было мозга, и создал пользователя, и привел его к программисту; и нарек программист его юзером. И сидели они оба под голым DOSом и не стыдились.
                        </Paragraph>
            </Col>
        </Row>
    )
}

export default Book