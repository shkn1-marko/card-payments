# 3. The Payment

An online card payment passes through a series of steps, from the moment the customer pays to the moment the merchant receives the money. At checkout, the cardholder enters the card details or confirms a payment with a digital wallet. The issuer first checks that the real cardholder is paying, then approves the payment and holds the amount, either in the cardholder's account or against their credit limit. Once the merchant captures the payment, the card network clears and settles it between the banks, and the acquirer pays the merchant. Even then, the payment can be reversed, by a refund from the merchant or a dispute raised by the cardholder.

## Checkout

The payment starts at **checkout**, the last step of an online purchase, where the customer confirms the order and chooses how to pay. To pay by card, the cardholder enters the card number, the expiry date and the security code from the back of the card. With a digital wallet or a card saved from an earlier purchase, the cardholder only confirms the payment, and a token is sent in place of the card details.

The payment form is often provided by the payment gateway, even when it appears on the merchant's own page. The card details go straight to the gateway and never pass through the merchant's systems, which spares the merchant most of the rules on protecting card data. The merchant supplies the amount, the currency and a reference to the order. The reference lets the merchant match the payment to the order later on.

The form catches simple mistakes, such as a mistyped card number, before the details are sent. Many payment service providers also screen each payment for signs of fraud, such as many attempts with different cards in a short time, and can stop it before it reaches any bank. If the customer leaves the checkout without confirming, the payment ends here, with no money moved and no bank involved.

## Authentication

Before the issuer approves a payment, it may need to check that the real cardholder is paying. This check is called **authentication**. In the EU, online card payments require **strong customer authentication**, which combines two of three kinds of proof: something the cardholder knows, such as a PIN or password, something they have, such as their phone, and something they are, such as a fingerprint or face.

The check runs through **3-D Secure**. The gateway sends the payment details through the card network to the issuer, together with information about the order and the cardholder's device. The issuer then decides whether to ask the cardholder to confirm. If it does, the cardholder approves the payment in their banking app, or enters a password and a code sent to their phone.

Not every payment needs the cardholder to act. Small payments up to €30 can be exempt, and so can payments the issuer or the acquirer judges safe from the information sent with the request. Regular payments of a subscription, after the cardholder has confirmed the first one, need no authentication at all. A payment made with a digital wallet often counts as authenticated already, since the cardholder has confirmed it on their phone.

When authentication succeeds, the issuer returns proof of it, which travels with the payment into authorization. The proof also protects the merchant: if an authenticated payment later turns out to be fraud, the loss in most cases falls on the issuer instead. If the cardholder fails to confirm, or does not respond in time, the payment stops here.

## Authorization

With the cardholder confirmed, the gateway asks the issuer to approve the payment itself. This request is called **authorization**. Along with the card details and the amount, it carries the proof the issuer returned during authentication.

The request passes from the gateway to the acquirer, or the processor working for it, and on to the card network. If the payment carries a token, the network or the payment service provider first swaps it back for the card number, whose first digits show which issuer the card belongs to. The network forwards the request to that issuer, or to its processor. If the issuer cannot be reached, the network can answer in its place, within limits the issuer has set in advance.

The issuer checks that the card is valid and not blocked, that the security code and expiry date match, and that the account holds enough money or credit. It also compares the payment with the cardholder's usual spending to spot fraud. The answer travels back the same way, within a few seconds.

An approval places a **hold** on the amount. The money stays in the cardholder's account but can no longer be spent, and on a credit card the available limit shrinks by the same amount. The cardholder may see it in their banking app as a pending payment. The merchant can now confirm the order, although no money has moved yet.

The issuer can also decline the payment, for reasons such as insufficient funds, an expired or blocked card, or suspected fraud. The issuer may give the exact reason for the decline, or just a general one. Merchants often show only a general message, and the customer can try again with another card.

## Capture

An approved payment only reserves the money. To collect it, the merchant must **capture** the payment, telling the acquirer to take the approved amount. The capture goes through the gateway to the acquirer, which adds the payment to those it will collect from the issuers.

Many merchants capture at once, together with the authorization, especially for digital goods and services delivered on the spot. Others wait. A shop that ships goods may capture when the order is sent, and a hotel may wait until the guest checks out. The merchant can capture less than was approved, for example when part of an order is out of stock. To take more, it needs a new approval from the issuer, since the hold covers only the original amount.

If the order is cancelled before capture, the merchant cancels the authorization instead. The issuer releases the hold, and the money is free to spend again without ever having left the cardholder's account. If the merchant neither captures nor cancels the payment, the hold expires on its own, typically after about a week.

## Clearing and Settlement

Captured payments reach the issuers through **clearing**. In a daily batch, the acquirer sends the payments its merchants have captured to the card network, which passes each one to its issuer. The issuer matches it to the authorization and replaces the hold with a final charge, taking the money from the cardholder's account or adding it to their credit card bill.

During clearing, the interchange fee is calculated for each payment. A payment in a foreign currency is converted at the exchange rate set at authorization, so the final charge should match the hold.

**Settlement** is the step in which the money moves, usually within one or two working days of capture. The issuer pays the acquirer the amount of each payment, minus the interchange fee, which the issuer keeps. Rather than moving money for each payment, the card network adds up what each bank owes and is owed for the day, and each bank pays or receives a single total. The network charges its scheme fees to both the issuer and the acquirer separately.

After settlement, the cardholder has been charged, and the acquirer has received the money.

## Payout

Once the acquirer has received the money, it pays the merchant. This is the **payout**. The acquirer adds up the merchant's payments, takes off the merchant service charge and any money returned to cardholders, and transfers the rest to the merchant's bank account. A payment facilitator does the same for the merchants signed up under it.

Payouts follow a schedule set in the merchant agreement, often daily or every few days. Some acquirers bill the merchant service charge once a month instead of taking it from each payout. A report with each payout shows how it was calculated, so the merchant can match the money received with its sales.

The acquirer or facilitator can also hold back part of the money for a time, as a safeguard against disputes it may later have to cover. This can follow from the risk that comes with the merchant's industry, or from unusual activity on its account.

With the payout, the payment is complete for every party.

## Refunds and Disputes

A completed payment can still be reversed. A **refund** is started by the merchant, for example when goods are returned or an order is cancelled after capture. The money travels back the way it came, through the acquirer and the card network to the issuer, and typically reaches the same card within one or two weeks. The merchant pays for it out of its next payout, and the fees on the original payment are often not returned.

A **dispute** is started by the cardholder. If they do not recognize a payment, or paid for goods that never arrived or were not as described, they can ask their issuer to get the money back, within a time limit that is most often 120 days. The issuer raises a **chargeback** through the card network. The money is taken back from the acquirer, which takes it from the merchant, often adding a fee.

The merchant can contest the chargeback by sending evidence through its acquirer, such as proof of delivery. The issuer reviews it and either returns the money to the merchant or keeps the chargeback in place. If the two sides still disagree, the card network makes the final decision. A merchant that receives many chargebacks can face penalties or lose the ability to accept cards.
