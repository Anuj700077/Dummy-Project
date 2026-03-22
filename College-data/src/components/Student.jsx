import React, { useState, useEffect } from 'react';
import '../css/Student.css';

function Student(){

  const [formData, setFormData] = useState({
    sname:"",
    fname:"",
    address:"",
    dob:""
  });

  const [students, setStudents] = useState([]);
  const [editId, setEditId] = useState(null);


  const handleChange = (e) =>{
    setFormData({...formData,[e.target.name]:e.target.value});
  }

  const fetchStudents = async () => {
  const response = await fetch("http://localhost:8080/students");
  const data = await response.json();

  console.log(data);  

  setStudents(data);
};


  useEffect(() => {
    fetchStudents();
  }, []);

const handleSubmit = async (e) => {
  e.preventDefault();

  if (editId === null) {

   
    await fetch("http://localhost:8080/students", {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify(formData)
    });

    alert("✅ Student added successfully");

  } else {

   
    await fetch(`http://localhost:8080/students/${editId}`, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify(formData)
    });

    alert("✅ Student updated successfully");

    setEditId(null);
  }

  setFormData({
    sname: "",
    fname: "",
    address: "",
    dob: ""
  });

  fetchStudents();
};

const handleEdit = (student) => {

  setFormData({
    sname: student.sname,
    fname: student.fname,
    address: student.address,
    dob: student.dob
  });

  setEditId(student.id);
};

    //for delete 

    const handleDelete = async (id) => {

  const confirmDelete = window.confirm("Are you sure you want to delete?");

  if (!confirmDelete) return;

  await fetch(`http://localhost:8080/students/${id}`, {
    method: "DELETE"
  });

  alert("Student deleted successfully");

  fetchStudents();
};



 // ends here

   return(
    <>
     <h1 style={{textAlign:"center"}}>Students Detail</h1>

     <div className='box1'>

     <div className='box2'>
      <form onSubmit={handleSubmit}>
      
        <label htmlFor="text">Student Name:</label>
        <input type="text" name='sname' placeholder='Enter you name'value={formData.sname} onChange={handleChange} /> <br />
        <br />
        <label htmlFor="text">Father Name:</label>
        <input type="text" name='fname' placeholder='Enter father name'value={formData.fname} onChange={handleChange}  /> <br />
        <br />
        <label htmlFor="text">Address:</label>
        <textarea name="address" id="address"value={formData.address} onChange={handleChange} ></textarea><br />
        <br />
        <label htmlFor="dob">Date of Birth:</label>
        <input type="date" name='dob' placeholder='Enter DOB'value={formData.dob} onChange={handleChange}  />
        <div className="button1">
          <button type='submit'>{editId === null ? "Submit" : "Update"}</button>
        </div>
      </form>
      </div>
</div>


 <h2 style={{textAlign:"center"}}>Students List</h2>

<div className='tableStyle'>
      <table border="1" cellPadding="10">
        <thead>
          <tr>
            <th>ID</th>
            <th>Student Name</th>
            <th>Father's Name</th>
            <th>Student's Address</th>
            <th>Student's DOB</th>
            <th>Action</th>
          </tr>
        </thead>

        <tbody>
          {students.length === 0 ? (
            <tr>
              <td colSpan="5" align="center">
                No data found
              </td>
            </tr>
          ) : (
            students.map((student) => (
              <tr key={student.id}>
                <td>{student.id}</td>
                <td>{student.sname}</td>
                <td>{student.fname}</td>
                <td>{student.address}</td>
                <td>{student.dob}</td>
                 <td>
                   <button onClick={() => handleEdit(student)}
                    style={{ backgroundColor: "lightgreen" }}>Edit</button>  

                     <button
                      onClick={() => handleDelete(student.id)}
                       style={{ backgroundColor: "red", color: "white" }}> Delete </button>
                    </td>
              </tr>
            ))
          )}
        </tbody>
      </table>

</div>

    </>
   )
}
export default Student;