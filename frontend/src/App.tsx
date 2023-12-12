import * as React from 'react';
import ToggleButton from '@mui/material/ToggleButton';
import ToggleButtonGroup from '@mui/material/ToggleButtonGroup';
import './App.css';
import {Toggle} from "../wailsjs/go/main/App";

function App() {
  const [alignment, setAlignment] = React.useState('web');

  const handleChange = (
    event: React.MouseEvent<HTMLElement>,
    newAlignment: string,
  ) => {
    setAlignment(newAlignment);
    Toggle(newAlignment);
  };

  return (
    <ToggleButtonGroup
      id="btngroup"
      color="primary"
      value={alignment}
      exclusive
      onChange={handleChange}
      aria-label="UpDown"
      fullWidth={true}
      size="large"
    >
      <ToggleButton color="success" value="up" aria-label="up">up</ToggleButton>
      <ToggleButton color="error" value="down" aria-label="down">down</ToggleButton>
    </ToggleButtonGroup>
  );
}

export default App
