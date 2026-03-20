import React from "react";
import "./App.css";
import { Link, Outlet } from "react-router-dom";

function App() {
  return (
    <div style={{ display: "flex" }}>
      
     
      <div className="sidebar">
        <h2>Menu</h2>
        <ul>
          <li><Link to="/student">Student</Link></li>
          <br />
          <li><Link to="/faculty">Faculty</Link></li>
          <br />
          <li><Link to="/marks">Students Mark</Link></li>
          <br />
          <li><Link to="/fees">Student Fees</Link></li>
        </ul>
      </div>

    
      <div style={{ flex: 1 }}>
        <h1 style={{ textAlign: "center" }}>Welcome to Dashboard</h1>

       
        <Outlet />
      </div>

    </div>
  );
}

export default App;
