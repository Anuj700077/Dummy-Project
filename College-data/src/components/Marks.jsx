import '../css/Marks.css'

function Marks(){
    return(
        <>
        <h1 style={{textAlign:"center"}}>Students Marks</h1>
        <div className="marks">
          <form className='mform'>
            <label htmlFor="text">Student Id:</label>
            <input type="text" name='sid' placeholder='Enter student id'/><br />
            <br />
            <label htmlFor="text">Subject</label>
            <select name='subject' id="">
                <option value="Mathematics">Mathematics</option>
                <option value="Science">Science</option>
                <option value="Hindi">Hindi</option>
                <option value="English">English</option>
                <option value="Computer">Computer</option>
            </select><br />
            <br />
            <label htmlFor="marks">Max Marks:</label>
            <input type="number" name='mm'/><br />
            <br />
            <label htmlFor="marks">Obtained Marks:</label>
            <input type="number" name='omark'/><br />
            <br />
             <label htmlFor="percentage">Percentage %:</label>
             <input type="percentage" name='percentage' /><br />
             <br />
             <label htmlFor="grade">Grade:</label>
             <select name='grade' id="">
                <option value="A">A</option>
                <option value="A+">A+</option>
                <option value="A++">A++</option>
                <option value="B">B</option>
                <option value="B+">B+</option>
                <option value="B++">B++</option>
                <option value="C">C</option>
                <option value="C+">C+</option>
                <option value="C++">C++</option>
                <option value="D">D</option>
             </select><br />
             <br />
                 <label htmlFor="ranks">Rank:</label>
                 <select name='rank' id="">
                    <option value="1">1</option>
                    <option value="2">2</option>
                    <option value="3">3</option>
                 </select><br />
                 <div className='btn3'>
                    <button type='submit'>Submit</button>
                 </div>
          </form>
        </div>
        </>
    )
}
export default Marks;