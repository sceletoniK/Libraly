import React, {useEffect, useState, useMemo} from 'react';
import axios, {AxiosError} from "axios";
import {useNavigate} from "react-router-dom";

const Refresh = () => {
    const navigate = useNavigate();

    axios.post(
        'http://localhost:8080/refresh',
        {},
        {
            headers: {
                Authorization: `Bearer ${localStorage.getItem('token')}`
            },
            withCredentials: true
        })
        .then((result) => {
            localStorage.setItem('token', result.data)
            navigate(-1);
        })
        .catch((error: AxiosError) => {
            navigate('/login')
        })

    return(
        <>
        </>
    )
}

export default  Refresh