import 'bootstrap/dist/css/bootstrap.css';
import defaultAvatar from '~/assets/images/avatar.jpg'
import './Users.css'
import { useEffect, useState} from "react";
import axios from "axios";
import {useParams} from "react-router-dom";

function DetailUser() {
    const {id} = useParams()
    console.log(id)
    const request = "http://localhost:8080/users/" + id
    console.log(request)
    const [user, setUser] = useState({})
    const getTokenFromLocalStorage = () => {
        return localStorage.getItem('token');
    };
    useEffect(() => {
      if (id) {
        axios.get(`http://localhost:8080/restricted/users/${id}`, {
            headers : {
                Authorization : `Bearer ${getTokenFromLocalStorage()}`,
            }
        }).then(res => {
            console.log(res.data)
            setUser(res.data)
        }).catch(error => {
            console.log('Error ', error)
        })
      }
    }, [])
    return (
        <div className="container">
            <div className="row" >
                <div className='col-4'>
                    <img src={defaultAvatar} className="img-fluid"/>
                </div>
                <div className='col-8'>
                    <div className='col-12 cus-header'>
                        USER INFOMATION
                    </div>
                    <div className='mt-1'>
                        <div className='cus-info'>First Name: {user?.first_name}</div>
                        <div className='cus-info'>Last Name: {user?.last_name}</div>
                        <div className='cus-info'>Address: {user?.address}</div>
                        <div className='cus-info'>Email: {user?.email}</div>
                        <div className='cus-info'>Phone Number: {user?.phone_number}</div>
                    </div>
                </div>
            </div>
        </div>
    )
}
export default DetailUser