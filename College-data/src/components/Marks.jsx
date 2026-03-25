import React, { useEffect, useState } from "react";
import "../css/Marks.css";

function Marks() {

  const [formData, setFormData] = useState({
    sid: "",
    math: "",
    science: "",
    hindi: "",
    english: "",
    computer: ""
  });

  const [students, setStudents] = useState([]);
  const [isEdit, setIsEdit] = useState(false); 

  
  const handleFormChange = (e) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  
  const handleSubmit = async (e) => {
    e.preventDefault();

    if (!formData.sid) {
      alert("Student ID required");
      return;
    }

    try {
      const url = "http://localhost:8080/marks";
      const method = isEdit ? "PUT" : "POST";

      const res = await fetch(url, {
        method: method,
        headers: {
          "Content-Type": "application/json"
        },
        body: JSON.stringify({
          sid: Number(formData.sid),
          math: Number(formData.math),
          science: Number(formData.science),
          hindi: Number(formData.hindi),
          english: Number(formData.english),
          computer: Number(formData.computer)
        })
      });

      const data = await res.json();

      if (!res.ok) {
        alert(data.error || "Something went wrong");
        return;
      }

      alert(data.message);

      fetchStudents();

      setFormData({
        sid: "",
        math: "",
        science: "",
        hindi: "",
        english: "",
        computer: ""
      });

      setIsEdit(false); 

    } catch (err) {
      console.error(err);
      alert("Error");
    }
  };

 
  const fetchStudents = async () => {
    try {
      const res = await fetch("http://localhost:8080/marks");
      const data = await res.json();
      setStudents(data);
    } catch (err) {
      console.error(err);
    }
  };

  useEffect(() => {
    fetchStudents();
  }, []);

 
  const handleEdit = (student) => {
    setFormData({
      sid: student.sid,
      math: student.math,
      science: student.science,
      hindi: student.hindi,
      english: student.english,
      computer: student.computer
    });

    setIsEdit(true);
  };

 
  const handleDelete = async (sid) => {

    if (!window.confirm("Are you sure to delete?")) return;

    try {
      const res = await fetch(`http://localhost:8080/marks/${sid}`, {
        method: "DELETE"
      });

      const data = await res.json();

      if (!res.ok) {
        alert(data.error);
        return;
      }

      alert(data.message);
      fetchStudents();

    } catch (err) {
      console.error(err);
      alert("Delete failed");
    }
  };

  return (
    <>
      <h1 style={{ textAlign: "center" }}>Students Marks</h1>

     
      <div className="marks">
        <form className="mform" onSubmit={handleSubmit}>

          <label>Student ID:</label>
          <input type="number" name="sid" value={formData.sid} onChange={handleFormChange} />
          <br /><br />

          <label>Math:</label>
          <input type="number" name="math" value={formData.math} onChange={handleFormChange} />
          <br /><br />

          <label>Science:</label>
          <input type="number" name="science" value={formData.science} onChange={handleFormChange} />
          <br /><br />

          <label>Hindi:</label>
          <input type="number" name="hindi" value={formData.hindi} onChange={handleFormChange} />
          <br /><br />

          <label>English:</label>
          <input type="number" name="english" value={formData.english} onChange={handleFormChange} />
          <br /><br />

          <label>Computer:</label>
          <input type="number" name="computer" value={formData.computer} onChange={handleFormChange} />
          <br /><br />

          <div className="btn3">
            <button type="submit">
              {isEdit ? "Update" : "Submit"}
            </button>
          </div>

        </form>
      </div>

      
      <h2 style={{ textAlign: "center" }}>Marks List</h2>

      <div className="marksEntry">
        <table border="1" cellPadding="10">
          <thead>
            <tr>
              <th>ID</th>
              <th>Name</th>
              <th>Math</th>
              <th>Science</th>
              <th>Hindi</th>
              <th>English</th>
              <th>Computer</th>
              <th>Total</th>
              <th>Percentage</th>
              <th>Action</th>
            </tr>
          </thead>

          <tbody>
            {students.length > 0 ? (
              students.map((student) => (
                <tr key={student.id}>
                  <td>{student.id}</td>
                  <td>{student.sname}</td>
                  <td>{student.math}</td>
                  <td>{student.science}</td>
                  <td>{student.hindi}</td>
                  <td>{student.english}</td>
                  <td>{student.computer}</td>
                  <td>{student.total}</td>
                  <td>{student.percentage}%</td>

                  <td>
                    <button onClick={() => handleEdit(student)}>Edit</button>
                    <button onClick={() => handleDelete(student.sid)}>Delete</button>
                  </td>
                </tr>
              ))
            ) : (
              <tr>
                <td colSpan="10" style={{ textAlign: "center" }}>
                  No data Found
                </td>
              </tr>
            )}
          </tbody>

        </table>
      </div>
    </>
  );
}

export default Marks;
