
import '../css/Fees.css'

function Fees(){
    return(
        <>
        <h1 style={{textAlign:"center"}}>Fees Details</h1>
        <div className='fees'>
            <form className='feeForm'>
                <label htmlFor="name">Student Name</label>
                <input type="name" name='sname' placeholder='Enter name' /> <br />
                <br />
                <label htmlFor="fname">Father name:</label>
                <input htmlFor="fname" name='fname' placeholder='Enter father name'></input><br />
                <br />
                <label htmlFor="text">Student Id:</label>
                <input htmlFor="number" name='sid' placeholder='Enter student id'></input><br />
                <br />
                <label htmlFor="text">Fee for Month:</label>
                <input type="date" name='feemonth' /><br />
                <br />
                <label htmlFor="number">Amount Paid:</label>
                <input type="number" name='amtpaid' placeholder='123456789' /> <br />
                <br />
                <label htmlFor="number">Amount Due:</label>
                <input type="number" name='amtdue' placeholder='123456789' /><br />
                <br />
                 <label htmlFor="text">Receivd Date:</label>
                <input type="date" name='receivedate' /><br />
                <br />
                <div className='btn1'>
                    <button type='submit'>Submit</button>
                </div>
            </form>
        </div>
        </>
    )
}
export default Fees;