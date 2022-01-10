import { LocalizationProvider, StaticDatePicker } from '@mui/lab'
import AdapterDateFns from '@mui/lab/AdapterDateFns'
import { Grid, TextField, Typography } from '@mui/material'
import { useState } from 'react'

export function FilterCalendar() {
  const [date, setValue] = useState<Date>(new Date())

  return (
    <Grid container item>
      <Typography>{date.toString()}</Typography>
      <LocalizationProvider dateAdapter={AdapterDateFns}>
        <StaticDatePicker
          displayStaticWrapperAs="desktop"
          value={null}
          onChange={(newValue) => {
            console.log(JSON.stringify(newValue))
            setValue(new Date(JSON.stringify(newValue)))
          }}
          renderInput={(params) => <TextField {...params} />}
        />
      </LocalizationProvider>
    </Grid>
  )
}
