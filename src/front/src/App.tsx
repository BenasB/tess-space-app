import { useState } from "react";
import "./App.css"
import LeftPane from './LeftPane';
import RightPane from './RightPane';

function App() {
  const [sector, setSector] = useState<number>(1)
  const [camera, setCamera] = useState<number>(2)
  const [ccd, setCcd] = useState<number>(1)

  return (
    <div className="split-screen-container">
      <div className="left-pane">
        <LeftPane sector={sector} camera={camera} ccd={ccd} setSector={setSector} setCamera={setCamera} setCcd={setCcd}></LeftPane>
      </div>
      <div className="right-pane">
        <RightPane sector={sector} camera={camera} ccd={ccd}></RightPane>
      </div>
    </div>
  )
}

export default App
