import { useState, useEffect } from "react";
import '../css/Fees.css';

function Fees() {

    const [formData, setFormData] = useState({
        sid: "",
        feemonth: "",
        amtpaid: "",
        receivedate: ""
    });

    const [history, setHistory] = useState([]);

  
    useEffect(() => {
        fetchAllFees();
    }, []);

    const fetchAllFees = async () => {
        try {
            const res = await fetch("http://localhost:8080/fees");
            const data = await res.json();

            if (Array.isArray(data)) {
                setHistory(data);
            } else {
                setHistory([]);
            }

        } catch {
            alert("Failed to fetch data");
        }
    };

    const handleChange = (e) => {
        setFormData({
            ...formData,
            [e.target.name]: e.target.value
        });
    };

    const handleSubmit = async (e) => {
        e.preventDefault();

        const payload = {
            sid: Number(formData.sid),
            feemonth: formData.feemonth,
            amtpaid: Number(formData.amtpaid),
            receivedate: formData.receivedate
        };

        try {
            const res = await fetch("http://localhost:8080/fees", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify(payload)
            });

            const data = await res.json();
            alert(data.message);

            if (res.ok) {
                fetchAllFees(); 
                setFormData({
                    sid: "",
                    feemonth: "",
                    amtpaid: "",
                    receivedate: ""
                });
            }

        } catch {
            alert(" Error submitting fee");
        }
    };

    return (
        <>
            <h1 style={{ textAlign: "center" }}>Students Fees</h1>

            <div className='fees'>
                <form className='feeForm' onSubmit={handleSubmit}>

                    <label>Student Id:</label>
                    <input type="number" name="sid" value={formData.sid} onChange={handleChange} /><br /><br />

                    <label>Fee Month:</label>
                    <input type="date" name="feemonth" value={formData.feemonth} onChange={handleChange} /><br /><br />

                    <label>Amount Paid:</label>
                    <input type="number" name="amtpaid" value={formData.amtpaid} onChange={handleChange} /><br /><br />

                    <label>Received Date:</label>
                    <input type="date" name="receivedate" value={formData.receivedate} onChange={handleChange} /><br /><br />

                    <button style={{marginLeft:"130px", marginTop:"40px"}} type="submit">Submit</button>
                </form>
            </div>

            <h2 style={{ textAlign: "center" }}>All Students Fees</h2>

            {history.length > 0 ? (
                <table border="1" style={{ margin: "auto" }} cellPadding='10'>
    <thead>
        <tr>
            <th>Student ID</th>
            <th>Student Name</th>
            <th>Father Name</th>
            <th>Month</th>
            <th>Paid</th>
            <th>Due</th>
            <th>Received Date</th>

        </tr>
    </thead>
    <tbody>
        {history.map((f, index) => (
            <tr key={index}>
                <td>{f.sid}</td>
                <td>{f.sname}</td>   
                <td>{f.fname}</td>   
                <td>{f.feemonth ? `${new Date(f.feemonth).toLocaleString
                ("en-US", { month: "long" })} - ${new Date(f.feemonth).getFullYear()}`: ""}</td>
                <td>{f.amtpaid}</td>
                <td>{f.amtdue}</td>
                <td> {f.receivedate ? new Date(f.receivedate).toLocaleDateString
                ("en-GB").split("/").join("-") : ""}</td>
            </tr>
        ))}
    </tbody>
</table>

            ) : (
                <p style={{ textAlign: "center" }}>No data available</p>
            )}
        </>
    );
}

export default Fees;
