import React, { useState, useEffect, useRef } from 'react'
import type { ColumnsType, ColumnType } from 'antd/es/table';
import type { InputRef } from 'antd';
import type { FilterConfirmProps } from 'antd/es/table/interface';
import { Button, Input, Space, Table, Col, Row } from 'antd';
import { SearchOutlined } from '@ant-design/icons';

const Books: React.FC = () => {

  const [error, setError] = useState<Error|null>(null);
  const [isLoaded, setIsLoaded] = useState(false);
  const [items, setItems] = useState<DataType[]>([]);

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
      )
  }, [])

  interface DataType {
    id: React.Key,
    name : string,
    author : string,
    publisher : string
  }

  type DataIndex = keyof DataType;

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

  const getColumnSearchProps = (dataIndex: DataIndex): ColumnType<DataType> => ({
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
    }
  });

  const columns : ColumnsType<DataType> = [
    {
      title: 'Название',
      dataIndex: 'name',
      key: 'name',
      ...getColumnSearchProps('name'),
    },
    
    {
      title: 'Автор',
      dataIndex: 'author',
      key: 'author',
      ...getColumnSearchProps('author'),
    },
    {
      title: 'Издатель',
      dataIndex: 'publisher',
      key: 'publisher',
      ...getColumnSearchProps('publisher'),
    },
  ];

  if (error) {
    return <div>Ошибка: {error.message}</div>;
  } else if (!isLoaded) {
    return <div>Загрузка...</div>;
  } else {
  return (
    <>
    <Row justify='center'>
      <Col span={18}>
        <Table bordered dataSource={items} columns={columns} />
      </Col>
    </Row>
    </>
  );
  }
}
  
export default Books;