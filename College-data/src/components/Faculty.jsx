import '../css/Faculty.css'

function Faculty(){
    return(
        <>
        <h1 style={{textAlign:"center"}}>About Faculty</h1>
     <div className="facultyForm">
       <form className='fForm'>
          <label htmlFor="name">Teacher Name:</label>
          <input type="text" name='tname' placeholder='Enter name' /> <br />
          <br />
          
          <label htmlFor="subject">Subject:</label>
           <select name='subject' id="">
           
          <option value="Mathematics">Mathematics</option>
          <option value="Science">Science</option>
          <option value="Hindi">Hindi</option>
          <option value="Computer">Computer</option>
          <option value="English">English</option>
          </select> <br />
    <br />
          <label htmlFor="text">Department:
            <select name='department' id="">
                <option value="CS">CS</option>
                <option value="IT">IT</option>
                <option value="ME">ME</option>
            </select><br />
         <br />
            <label htmlFor="date">Date of Assign:</label>
            <input type="date" name='doa'  />
          </label> <br />
           <div className='btn'>
            <button type='submit'>Submit</button>
           </div>
       </form>
      </div>

        </>
    )
}
export default Faculty;