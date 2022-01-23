import { LocalizationProvider, StaticDatePicker } from '@mui/lab'
import AdapterDateFns from '@mui/lab/AdapterDateFns'
import { Grid, TextField, Typography } from '@mui/material'
import { useState } from 'react'

export function FilterCalendar() {
  const [date, setDate] = useState<Date>(new Date())

  return (
    <Grid container item>
      <Typography>{date && date.toString()}</Typography>
      <LocalizationProvider dateAdapter={AdapterDateFns}>
        <StaticDatePicker
          disablePast={true}
          displayStaticWrapperAs="desktop"
          value={date}
          onChange={(newValue) => {
            setDate(newValue || date)
          }}
          renderInput={(params) => <TextField {...params} />}
        />
      </LocalizationProvider>
    </Grid>
  )
}

// import TextField from '@mui/material/TextField'
// import StaticDateRangePicker from '@mui/lab/StaticDateRangePicker'
// import AdapterDateFns from '@mui/lab/AdapterDateFns'
// import LocalizationProvider from '@mui/lab/LocalizationProvider'
// import Box from '@mui/material/Box'
// import { DateRange } from '@mui/lab/DateRangePicker'
// import { useState } from 'react'

// export function FilterCalendar() {
//   const [value, setValue] = useState<DateRange<Date>>([null, null])

//   return (
//     <LocalizationProvider dateAdapter={AdapterDateFns}>
//       <StaticDateRangePicker
//         displayStaticWrapperAs="desktop"
//         disableAutoMonthSwitching={true}
//         disablePast={true}
//         value={value}
//         onChange={(newValue) => {
//           setValue(newValue)
//         }}
//         renderInput={(startProps, endProps) => (
//           <>
//             <TextField {...startProps} />
//             <Box sx={{ mx: 2 }}> to </Box>
//             <TextField {...endProps} />
//           </>
//         )}
//       />
//     </LocalizationProvider>
//   )
// }
