import { useState, useEffect } from 'react'
import '../css/Faculty.css'

function Faculty() {


const [formData, setFormData] = useState({
tname: '',
subject: '',
department: '',
doa: ''
})


const [facultyList, setFacultyList] = useState([])


const [editId, setEditId] = useState(null)


const handleChange = (e) => {
setFormData({
...formData,
[e.target.name]: e.target.value
})
}


const fetchFaculty = async () => {
try {
const res = await fetch("http://localhost:8080/faculty")
const data = await res.json()
setFacultyList(data)
} catch (err) {
console.log(err)
}
}


useEffect(() => {
fetchFaculty()
}, [])


const handleEdit = (f) => {
setFormData({
tname: f.tname,
subject: f.subject,
department: f.department,
doa: f.doa
})
setEditId(f.id)
}


const handleSubmit = async (e) => {
e.preventDefault()


try {
  let url = "http://localhost:8080/faculty"
  let method = "POST"

  if (editId) {
    url = `http://localhost:8080/faculty/${editId}`
    method = "PUT"
  }

  const res = await fetch(url, {
    method: method,
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify(formData)
  })

  const data = await res.json()
  console.log(data)

  alert(editId ? "Faculty updated successfully" : "Faculty added successfully")

 
  fetchFaculty()

 
  setFormData({
    tname: '',
    subject: '',
    department: '',
    doa: ''
  })

  setEditId(null)

} catch (err) {
  console.log(err)
}



}
const handleDelete = async (id) => {
  const confirmDelete = window.confirm("Are you sure you want to delete?")
  if (!confirmDelete) return

  try {
    const res = await fetch(`http://localhost:8080/faculty/${id}`, {
      method: "DELETE"
    })

    const data = await res.json()
    console.log(data)

    alert("Faculty deleted successfully")

    fetchFaculty()

  } catch (err) {
    console.log(err)
  }
}
return (
<>
<h1 style={{ textAlign: "center" }}>About Faculty</h1>


  <div className="facultyForm">
    <form className='fForm' onSubmit={handleSubmit}>

      <label>Teacher Name:</label>
      <input type="text" name='tname' value={formData.tname} onChange={handleChange} /><br />
      <br />

      <label>Subject:</label>
      <select name='subject' value={formData.subject} onChange={handleChange}>
        <option value="">Select</option>
        <option value="Mathematics">Mathematics</option>
        <option value="Science">Science</option>
        <option value="Hindi">Hindi</option>
        <option value="Computer">Computer</option>
        <option value="English">English</option>
      </select><br /><br />

      <label>Department:</label>
      <select name='department' value={formData.department} onChange={handleChange}>
        <option value="">Select</option>
        <option value="CS">CS</option>
        <option value="IT">IT</option>
        <option value="ME">ME</option>
      </select><br /><br />

      <label>Date of Assign:</label>
      <input type="date" name='doa' value={formData.doa} onChange={handleChange} /><br />
      <br />

      <div className='btn'>
        <button type='submit'>
          {editId ? "Update" : "Submit"}
        </button>
      </div>

    </form>
  </div>

 <h2 style={{ textAlign: "center" }}>Faculty List</h2>

  <table border="1" style={{ margin: "20px auto", width: "60%" }}>
    <thead>
      <tr>
        <th>ID</th>
        <th>Name</th>
        <th>Subject</th>
        <th>Department</th>
        <th>Date</th>
        <th>Action</th>
      </tr>
    </thead>

    <tbody>
      {facultyList.map((f) => (
        <tr key={f.id}>
          <td>{f.id}</td>
          <td>{f.tname}</td>
          <td>{f.subject}</td>
          <td>{f.department}</td>
          <td>{f.doa}</td>
          <td>
            <button onClick={() => handleEdit(f)}>Edit</button>
             <button onClick={() => handleDelete(f.id)}>Delete</button>
          </td>
        </tr>
      ))}
    </tbody>
  </table>
</>


)
}

export default Faculty;
