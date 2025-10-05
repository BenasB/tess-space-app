import { useState } from "react";
import "./App.css"
import LeftPane from './LeftPane';
import RightPane from './RightPane';
import { useMemo } from 'react';
import targetData from './targets.json';

type TargetData = {
  [sector: string]: {
    [camera: string]: {
      [ccd: string]: any[];
    };
  };
};

function App() {
  const [sector, setSector] = useState<number>(1)
  const [camera, setCamera] = useState<number>(1)
  const [ccd, setCcd] = useState<number>(1)

  const currentTargets =
    (targetData as TargetData)?.[sector]?.[camera]?.[ccd] || [];

  const sortedTargets = useMemo(() => {
    return currentTargets.sort((a, b) => b.Tmag - a.Tmag)
  }, [sector, camera, ccd])

  return (
    <div className="split-screen-container">
      <div className="left-pane">
        <LeftPane sector={sector} camera={camera} ccd={ccd} setSector={setSector} setCamera={setCamera} setCcd={setCcd} targets={sortedTargets}></LeftPane>
      </div>
      <div className="right-pane">
        <RightPane sector={sector} camera={camera} ccd={ccd} targets={sortedTargets}></RightPane>
      </div>
    </div>
  )
}

export default App
