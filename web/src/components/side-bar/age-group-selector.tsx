import {
  ToggleButtonGroup,
  ToggleButton,
  Grid,
  Typography,
} from '@mui/material'
import { useState } from 'react'

export function AgeGroupSelector() {
  const [alignment, setAlignment] = useState('')

  const handleChange = (
    event: React.MouseEvent<HTMLElement>,
    newAlignment: string
  ) => {
    setAlignment(newAlignment)
  }

  return (
    <Grid container item justifyContent="center">
      {/* <Typography fontWeight="bold">Age group</Typography>   */}
      <ToggleButtonGroup value={alignment} exclusive onChange={handleChange}>
        <ToggleButton value="0" color="success">
          L
        </ToggleButton>
        <ToggleButton value="10" color="primary">
          10
        </ToggleButton>
        <ToggleButton value="12" color="warning">
          12
        </ToggleButton>
        <ToggleButton value="14" color="error">
          14
        </ToggleButton>
        <ToggleButton value="16" color="info">
          16
        </ToggleButton>
        <ToggleButton value="18">18</ToggleButton>
      </ToggleButtonGroup>
    </Grid>
  )
}
