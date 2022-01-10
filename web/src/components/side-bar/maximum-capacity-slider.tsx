import { Grid, Slider, Typography } from '@mui/material'
import { useState } from 'react'

export function MaximumCapacitySlider() {
  const [capacity, setCapacity] = useState(0)
  return (
    <Grid container item>
      <Typography fontWeight="bold">
        Maximum capacity of people: {capacity}
      </Typography>
      <Slider
        aria-label="maximum-capacity"
        defaultValue={200}
        valueLabelDisplay="auto"
        min={10}
        max={300}
        onChangeCommitted={(e, v) => setCapacity(v as number)}
      />
    </Grid>
  )
}
