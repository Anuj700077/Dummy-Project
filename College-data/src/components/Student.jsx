import React from 'react';
import '../css/Student.css';
function Student(){
   return(
    <>
     <h1 style={{textAlign:"center"}}>Students Detail</h1>

     <div className='box1'>

     <div className='box2'>
      <form>
      
        <label htmlFor="text">Student Name:</label>
        <input type="text" name='sname' placeholder='Enter you name' /> <br />
        <br />
        <label htmlFor="text">Father Name:</label>
        <input type="text" name='fname' placeholder='Enter father name' /> <br />
        <br />
        <label htmlFor="text">Address:</label>
        <textarea name="address" id="address"></textarea><br />
        <br />
        <label htmlFor="dob">Date of Birth:</label>
        <input type="date" name='dob' placeholder='Enter DOB' />
        <div className="button1">
          <button style={{color:"blue"}}>Submit</button>
        </div>
      </form>
      </div>
</div>
    </>
   )
}
export default Student;