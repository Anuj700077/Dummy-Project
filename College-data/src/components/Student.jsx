import React, { useState } from 'react';
import '../css/Student.css';
function Student(){
  
   //writing these lines for connecting to backend
    const [formData, setFormData] = useState({
      sname:"",
      fname:"",
      address:"",
      dob:""
    });

    const handleChange = (e) =>{
      setFormData({...formData,[e.target.name]:e.target.value});
    }

   const handleSubmit = async (e) =>{
    e.preventDefault();

    await fetch("http://localhost:8080/students",{
      method:"POST",
      headers:{
        "Content-Type":"application/json"
      },
      body: JSON.stringify(formData)
    });
    alert(" ✅✅Student added Successfully ✅✅")
   }
 // ends here

   return(
    <>
     <h1 style={{textAlign:"center"}}>Students Detail</h1>

     <div className='box1'>

     <div className='box2'>
      <form onSubmit={handleSubmit}>
      
        <label htmlFor="text">Student Name:</label>
        <input type="text" name='sname' placeholder='Enter you name' onChange={handleChange} /> <br />
        <br />
        <label htmlFor="text">Father Name:</label>
        <input type="text" name='fname' placeholder='Enter father name' onChange={handleChange}  /> <br />
        <br />
        <label htmlFor="text">Address:</label>
        <textarea name="address" id="address" onChange={handleChange} ></textarea><br />
        <br />
        <label htmlFor="dob">Date of Birth:</label>
        <input type="date" name='dob' placeholder='Enter DOB' onChange={handleChange}  />
        <div className="button1">
          <button type='submit' style={{color:"blue"}}>Submit</button>
        </div>
      </form>
      </div>
</div>
    </>
   )
}
export default Student;