import 'dotenv/config'
import express from 'express'
import { env } from 'process'
import Stripe from 'stripe'

const app = express()

// connect to stripe
const client = new Stripe(env.STRIPE_API_KEY, {
  apiVersion: '2020-08-27',
  typescript: true,
})

app.get('/connect', async (req, res) => {
  // create an account
  const params: Stripe.AccountCreateParams = {
    type: 'express',
    country: 'BR',
    email: 'test@gmail.com',
    default_currency: 'BRL',
    business_type: 'individual',
    individual: {
      email: 'gluislopes011@gmail.com',
    },
    settings: {
      payouts: { schedule: { delay_days: 'minimum' } },
    },
    capabilities: {
      transfers: { requested: true },
      card_payments: { requested: true },
    },
  }
  const account = await client.accounts.create(params)

  // create account link
  const accountLink = await client.accountLinks.create({
    account: account.id,
    refresh_url: 'http://localhost:4040/connect',
    return_url: 'http://localhost:4040/connected',
    type: 'account_onboarding',
  })
  console.log('new account link created')
  console.log(accountLink)

  res.redirect(accountLink.url)
})

app.get('/payment', async (req, res) => {
  const id = req.query['id'] as string
  // const charge = await client.paymentIntents.create({
  //   amount: 1000,
  //   currency: 'brl',
  //   payment_method: 'pm_card_visa',
  //   transfer_data: {
  //     destination: id,
  //   },
  // })

  // create checkout session
  // const charge = await client.charges.create({
  //   amount: 100,
  //   application_fee_amount: Number(env.MORPHEUS_FEE_PERCENTAGE),
  //   currency: 'BRL',
  //   source: '',
  //   destination: {
  //     account: id,
  //   },
  // })
  const session = await client.checkout.sessions.create({
    line_items: [
      {
        price_data: {
          currency: 'brl',
          product_data: {
            name: 'Ticket Enrique e Juliano',
          },
          unit_amount: 2000,
        },
        quantity: 3,
      },
    ],
    mode: 'payment',
    success_url: 'https://example.com/success',
    cancel_url: 'https://example.com/failure',
    payment_intent_data: {
      application_fee_amount: 123,
      on_behalf_of: id,
      transfer_data: {
        destination: id,
      },
    },
  })

  res.redirect(session.url)
  // res.json(charge)
})

app.get('/payout', async (req, res) => {
  const id = req.query['id'] as string

  const payout = await client.payouts.create(
    {
      amount: 1000,
      currency: 'brl',
    },
    {
      stripeAccount: id,
    }
  )

  res.json(payout)
})

app.get('/connected', (req, res) => {
  res.json({ message: 'connected successfully' })
})

app.listen(4040)

// async function createConnectedClient(): Promise<void> {
//   try {
//     const client = connectStripe()
//     //create account
//     const params: Stripe.AccountCreateParams = {
//       type: 'custom',
//       country: 'BR',
//       email: 'gluislopes011@gmail.com',
//       default_currency: 'BRL',
//       business_type: 'individual',
//       individual: {
//         email: 'gluislopes011@gmail.com',
//       },
//       settings: {
//         payouts: { schedule: { delay_days: 'minimum' } },
//       },
//       capabilities: {
//         transfers: { requested: true },
//         card_payments: { requested: true },
//       },
//     }
//     const account = await client.accounts.create(params)
//     // account.charges_enabled = true
//     // account.payouts_enabled = true
//     console.log('new account created')
//     // console.log(account)

//     // create account link
//     const accountLink = await client.accountLinks.create({
//       account: account.id,
//       refresh_url: 'https://example.com/reauth',
//       return_url: 'https://example.com/return',
//       type: 'account_onboarding',
//     })
//     console.log('new account link created')
//     console.log(accountLink)

//     // create checkout session
//     // const charge = await client.charges.create({
//     //   amount: 100,
//     //   application_fee_amount: Number(env.MORPHEUS_FEE_PERCENTAGE),
//     //   currency: 'BRL',
//     //   // customer: 'acct_1KpzX9B7FXZVU1Uc',
//     //   source: '',
//     //   destination: {
//     //     account: account.id,
//     //   },
//     // })

//     // const paymentIntent = await client.paymentIntents.create({
//     //   amount: 1000,
//     //   currency: 'brl',
//     //   application_fee_amount: 123,
//     //   transfer_data: {
//     //     destination: account.id,
//     //   },
//     // })

//     // const session = await client.checkout.sessions.create({
//     //   line_items: [
//     //     {
//     //       price: 'price_1Kpz6HBC0LpbEEZfwRm8UYp9',
//     //       quantity: 1,
//     //     },
//     //   ],
//     //   mode: 'payment',
//     //   success_url: 'https://example.com/success',
//     //   cancel_url: 'https://example.com/failure',
//     //   payment_intent_data: {
//     //     application_fee_amount: 123,
//     //     transfer_data: {
//     //       // destination: '{{CONNECTED_ACCOUNT_ID}}',
//     //       destination: account.id,
//     //     },
//     //   },
//     // })
//     // console.log('new payment created')
//     // console.log(session)
//   } catch (error) {
//     console.log('error while creating an account')
//     console.log(error)
//   }
// }

// ;(async () => {
//   console.log('here')

//   await createConnectedClient()
// })()
