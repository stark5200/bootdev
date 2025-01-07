

import React, { useEffect, useState } from 'react'
import { AiOutlineSearch } from "react-icons/ai";
import UserLoginListElement from './UserLoginListElement';
import PasswordField from './PasswordField';
import { GoDecrypt } from "../../wailsjs/go/backend/App";
import Users from '../../../users.json'

import './UserLoginsList.css'

function UserLoginsList()
{
  //let loggedIn = false
  let falsePassword = "pseudoPassword123"
  const [passw, setPassw] = useState(falsePassword)
  const [login, setLogin] = useState(true)

  setPassw

  useEffect(() => { 
    console.log("loged in is: ", login)
  }, [login]);

  interface User {
    name: string;
    username: string;
    encryptedPassword: string;
    privateKey: string;
    categories: string[];
  }

type UsersType = {
  [key: string]: User;
};

  const typedUsers = Users as unknown as UsersType;

  // Get the keys of the Users object
  const userKeys = Object.keys(typedUsers);
  //<PasswordField password={loggedIn ? GoDecrypt(user.encryptedPassword, user.privateKey) : falsePassword} />

  return (
    <div id="UserLoginList">    
      {userKeys.map((key, index) => {
        const user = typedUsers[key];
        return (
          <div key={index}>
            <UserLoginListElement
              name={user.name}
              username={user.username}
              categories={user.categories}
            />
            <PasswordField password={passw} />
            <button onClick={() => { setLogin(!login),  setPassw(GoDecrypt) }}>{ String(login) setPassw()}</button>
            <hr />
          </div>
        );
      })}
      <button className="btn btn-info sticky-btn" id="sticky-btn-search"><AiOutlineSearch /></button>
      <button className="btn btn-info sticky-btn" id="sticky-btn-add">+</button>

    </div>

  )

}

export default UserLoginsList
