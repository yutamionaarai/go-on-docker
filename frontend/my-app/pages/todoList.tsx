import React, { useState, useEffect } from 'react'
import axios from 'axios';
import type { Todo } from '../../my-app/types/todo';


const TodoList = () => {
    const [data, setData] = useState([])
    useEffect(() => {
        const getAllTodos = async () => {
        const res = await axios.get('http://localhost:8080/todos/12');
        console.log(1245)
        console.log(res.data[0])
        setData(res.data)
        }
        getAllTodos()
      }, [])

      return (
        <>
        <p>YouTube</p>
          {/* {data.map((todos) => {
            const { id, title, description } = todos
            return (
              <div key={id}>
                <p>{title}</p>
              </div>
            );
          })} */}
        </>
      );
    }






//    {data.map((todo) => {
//         const { id, title, description } = todo
//    }

//    {data.map((todo) => {
//     const { id, title, description } = todo

//     return (
//     )
   
// }




export default TodoList