import React, { useState } from 'react';
import { useLocation } from 'react-router-dom';
import {
  DesktopOutlined,
  FileOutlined,
  PieChartOutlined,
  TeamOutlined,
  UserOutlined,
} from '@ant-design/icons';
import {Breadcrumb, Divider, Layout, Menu, theme, Typography} from 'antd';
import { Routes, Route, Link } from 'react-router-dom';
import Books from './pages/Books/Books';
import Login from './pages/Login/Login';
import Registration from './pages/Registration/Registation';
import Account from "./pages/Account/Account";
import Refresh from "./pages/Refresh/Refresh";
import Book from "./pages/Book/Book";

const { Title, Text } = Typography;

const { Header, Content, Footer, Sider } = Layout;

const App: React.FC = () => {
  const [collapsed, setCollapsed] = useState(false);
  const location = useLocation();
  const {
    token: { colorBgContainer },
  } = theme.useToken();

  return (
    <Layout style={{ minHeight: '100vh' }}>
      <Sider collapsible collapsed={collapsed} onCollapse={(value) => setCollapsed(value)}>
        <Menu
            theme="dark"
            defaultSelectedKeys={['/']}
            selectedKeys={[ location.pathname ]}
            mode="inline"
        >
        <div style={{ height: 32, margin: 16, background: 'rgba(255, 255, 255, 0.2)' }} />
          <Menu.Item key='/' icon={<FileOutlined />}>
            <Link to='/'>Ассортимент</Link>
          </Menu.Item>
          <Menu.Item key='/account' icon={<PieChartOutlined />}>
            <Link to='/account'>Личный кабинет</Link>
          </Menu.Item>
        </Menu>
      </Sider>
      <Layout className="site-layout">
        <Header style={{ padding: 0, background: colorBgContainer }}>
          <Title style={{ margin: 8, textAlign: "center"}}>Библиотека</Title>
        </Header>
        <Content style={{ margin: '0 16px' }}>
          <Divider dashed/>
          <div style={{ padding: 24, minHeight: '100%', background: colorBgContainer }}>
            <Routes>
              <Route path='/' element={<Books />} />
              <Route path='/account' element={<Account />}/>
              <Route path='/book/:id' element={<Book />}/>
              <Route path='/reg' element={<Registration />}/>
              <Route path='/login' element={<Login />}/>
              <Route path='/refresh' element={<Refresh/>}/>
            </Routes>
          </div>
        </Content>
        <Footer style={{ textAlign: 'center' }}>Ant Design ©2023 Created by Ant UED</Footer>
      </Layout>
    </Layout>
  );
};

export default App;