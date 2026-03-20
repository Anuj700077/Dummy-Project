import { BrowserRouter, Routes, Route } from 'react-router-dom';
import App from './App';
import Student from './components/Student';
import Faculty from './components/Faculty';
import Fees from './components/Fees';
import Marks from './components/Marks';

function MainRoute() {
  return (
    <BrowserRouter>
      <Routes>

        {/* Parent Route */}
        <Route path="/" element={<App />}>
          <Route path="student" element={<Student />} />
          <Route path="faculty" element={<Faculty />} />
          <Route path="fees" element={<Fees />} />
          <Route path="marks" element={<Marks />} />
        </Route>

      </Routes>
    </BrowserRouter>
  );
}

export default MainRoute;
