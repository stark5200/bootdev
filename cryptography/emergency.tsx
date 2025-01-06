import React from 'react'
import { AiOutlineSearch } from "react-icons/ai";
import UserLoginListElement from './UserLoginListElement';
import PasswordField from './PasswordField';
import Users from '../../../users.json'

import './UserLoginsList.css'

function UserLoginsList()
{
  let loggedIn = false
  let falsePassword = "pseudoPassword123"

  interface User {
    Name: string;
    Username: string;
    EncryptedPassword: string;
    PrivateKey: string;
    Categories: string[];
  }

type UsersType = {
  [key: string]: User;
};

  const typedUsers = Users as unknown as UsersType;

  // Get the keys of the Users object
  const userKeys = Object.keys(typedUsers);

  return (
    <div id="UserLoginList">    
      {userKeys.map((key, index) => {
        const user = typedUsers[key];
        return (
          <div key={index}>
            <UserLoginListElement
              name={user.Name}
              username={user.Username}
              categories={user.Categories}
            />
            <PasswordField password={loggedIn ? user.EncryptedPassword : falsePassword} />
          </div>
        );
      })}
      <button className="btn btn-info sticky-btn" id="sticky-btn-search"><AiOutlineSearch /></button>
      <button className="btn btn-info sticky-btn" id="sticky-btn-add">+</button>

    </div>

  )

}

export default UserLoginsList

/*

import React from 'react'
import { AiOutlineSearch } from "react-icons/ai";
import UserLoginListElement from './UserLoginListElement';
import PasswordField from './PasswordField';
import Users from '../../../users.json'

import './UserLoginsList.css'

function UserLoginsList()
{
  let loggedIn = false
  let falsePassword = "pseudoPassword123"

  interface User {
    Name: string;
    Username: string;
    EncryptedPassword: string;
    PrivateKey: string;
    Categories: string[];
  }

type UsersType = {
  [key: string]: User;
};

  const typedUsers = Users as unknown as UsersType;

  // Get the keys of the Users object
  const userKeys = Object.keys(typedUsers);

    return (

        <div id="UserLoginList">
      {userKeys.map((key, index) => {
        const user = typedUsers[key];
        return (
          <div key={index}>
            <UserLoginListElement
              name={user.Name}
              username={user.Username}
              categories={user.Categories}
            />
            <PasswordField password={loggedIn ? user.EncryptedPassword : falsePassword} />
          </div>
        );
      })}
            <button className="btn btn-info sticky-btn" id="sticky-btn-search"><AiOutlineSearch /></button>
            <button className="btn btn-info sticky-btn" id="sticky-btn-add">+</button>

        </div>

    )

}

export default UserLoginsList

*/