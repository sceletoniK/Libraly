import React, { useState } from 'react';
import {
  DesktopOutlined,
  FileOutlined,
  PieChartOutlined,
  TeamOutlined,
  UserOutlined,
} from '@ant-design/icons';
import { Breadcrumb, Layout, Menu, theme } from 'antd';
import { Routes, Route, Link } from 'react-router-dom';
import Books from './pages/Books';
import Account from './pages/Login';

const { Header, Content, Footer, Sider } = Layout;

const App: React.FC = () => {
  const [collapsed, setCollapsed] = useState(false);
  const {
    token: { colorBgContainer },
  } = theme.useToken();

  return (
    <Layout style={{ minHeight: '100vh' }}>
      <Sider collapsible collapsed={collapsed} onCollapse={(value) => setCollapsed(value)}>
        <Menu theme="dark" defaultSelectedKeys={[ window.location.pathname ]} mode="inline">
        <div style={{ height: 32, margin: 16, background: 'rgba(255, 255, 255, 0.2)' }} />
          <Menu.Item key='/'icon={<FileOutlined />}>
            <Link to='/'>Ассортимент</Link>
          </Menu.Item>
          <Menu.Item key='/account' icon={<PieChartOutlined />}>
            <Link to='/account'>Личный кабинет</Link>
          </Menu.Item>
        </Menu>
      </Sider>
      <Layout className="site-layout">
        <Header style={{ padding: 0, background: colorBgContainer }}>
          <p className='header_text'>Библиотечка</p>
        </Header>
        <Content style={{ margin: '0 16px' }}>
          <Breadcrumb style={{ margin: '16px 0' }}>
            <Breadcrumb.Item>Books</Breadcrumb.Item>
          </Breadcrumb>
          <div style={{ padding: 24, minHeight: '100%', background: colorBgContainer }}>
            <Routes>
              <Route path='/' element={<Books />} />
              <Route path='/account' element={<Account />}/>
            </Routes>
          </div>
        </Content>
        <Footer style={{ textAlign: 'center' }}>Ant Design ©2023 Created by Ant UED</Footer>
      </Layout>
    </Layout>
  );
};

export default App;