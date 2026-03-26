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
    const [isHistoryView, setIsHistoryView] = useState(false);

    useEffect(() => {
        fetchFees();
    }, []);

   
    const fetchFees = async () => {
        try {
            const res = await fetch("http://localhost:8080/fees");
            const data = await res.json();

            if (Array.isArray(data)) {
                setHistory(data);
            } else {
                console.error("Invalid response:", data);
                setHistory([]);
            }

            setIsHistoryView(false);

        } catch (err) {
            console.error(err);
            alert("Failed to fetch data");
        }
    };

    
    const fetchStudentHistory = async (sid) => {
        try {
            const res = await fetch(`http://localhost:8080/fees/student/${sid}`);
            const data = await res.json();

            if (Array.isArray(data)) {
                setHistory(data);
                setIsHistoryView(true);
            } else {
                console.error("Invalid history response:", data);
                alert("No history found");
            }

        } catch (err) {
            console.error(err);
            alert("Failed to fetch student history");
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
                fetchFees();
                setFormData({
                    sid: "",
                    feemonth: "",
                    amtpaid: "",
                    receivedate: ""
                });
            }

        } catch (err) {
            console.error(err);
            alert("Error submitting fee");
        }
    };

   
    const formatMonth = (date) => {
        if (!date) return "N/A";
        const d = new Date(date);
        return isNaN(d) ? "N/A" :
            `${d.toLocaleString("en-US", { month: "long" })} - ${d.getFullYear()}`;
    };

    const formatDate = (date) => {
        if (!date) return "N/A";
        const d = new Date(date);
        return isNaN(d) ? "N/A" :
            d.toLocaleDateString("en-GB").split("/").join("-");
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

                    <button type="submit">Submit</button>
                </form>
            </div>

           
            {isHistoryView && (
                <div style={{ textAlign: "center", marginTop: "20px" }}>
                    <button onClick={fetchFees}>⬅ Back</button>
                </div>
            )}

            <h2 style={{ textAlign: "center" }}>
                {isHistoryView ? "Student Fee History" : "Students Fees"}
            </h2>

            {history.length > 0 ? (
                <table border="1" style={{ margin: "auto" }} cellPadding='10'>
                    <thead>
                        <tr>
                            <th>Student ID</th>
                            <th>Name</th>
                            <th>Father</th>
                            <th>Month</th>
                            <th>Paid</th>
                            <th>Due</th>
                            <th>Date</th>
                            {!isHistoryView && <th>Action</th>}
                        </tr>
                    </thead>

                    <tbody>
                        {history.map((f, index) => (
                            <tr key={index}>
                                <td>{f.sid}</td>
                                <td>{f.sname}</td>
                                <td>{f.fname}</td>

                                <td>{formatMonth(f.feemonth)}</td>

                                <td>{f.amtpaid}</td>
                                <td>{f.amtdue}</td>

                                <td>{formatDate(f.receivedate)}</td>

                                {!isHistoryView && (
                                    <td>
                                        <button onClick={() => fetchStudentHistory(f.sid)}>
                                            View History
                                        </button>
                                    </td>
                                )}
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
