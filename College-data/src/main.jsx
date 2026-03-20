import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import MainRoute from './MainRoute.jsx'

createRoot(document.getElementById('root')).render(
  <StrictMode>
    <MainRoute/>
  </StrictMode>,
)
