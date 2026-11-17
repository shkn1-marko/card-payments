# 2. The Players

Most card payments rely on a chain of companies. In the middle, a card network, most often Visa or Mastercard, connects the issuer, which gives the customer the card, with the acquirer, which accepts payments for the merchant. Around them, processors and gateways carry out the technical work, payment facilitators bring small businesses into the system, and digital wallets let customers pay without typing their card details. The roles are divided for reasons of reach, risk and cost. The card networks hold the most power, while merchants have more choice among the companies below them. Every payment carries a fee, paid by the merchant and shared among the players. In Croatia, most players are owned by foreign groups.

## The Four-Party Model

Most card payments involve four parties. On one side are the **cardholder**, who pays, and the **issuing bank**, or issuer, which gave them the card. On the other are the **merchant**, the shop that gets paid, and the **acquiring bank**, or acquirer, which accepts card payments on the merchant's behalf. Between the two sides sits the **card network**, such as Visa or Mastercard, which sets the rules and carries messages between issuers and acquirers. The network connects the parties without being one of them, which is why the arrangement is called the **four-party model**.

When a customer pays online, the request travels from the merchant to the acquirer, through the card network, to the issuer. The issuer checks the card and the available funds, then approves or declines the payment. The answer travels back the same way, usually within seconds. The money follows later: the issuer pays the acquirer through the network, and the acquirer pays the merchant, minus fees.

The strength of the model is that the two sides never need to know each other. A card from any issuer works at any merchant, as long as the issuer and the merchant's acquirer belong to the same network.

Online, more companies stand between the merchant and the acquirer, such as processors, gateways and payment facilitators. They are covered later in this chapter, and they leave the model unchanged: each one works on behalf of one of the four parties.

## Issuing Banks

The issuing bank, or issuer, represents the cardholder in the payment. It gives the customer the card, links it to an account, and decides whether each payment goes through.

Cards come in three main types, depending on where the money comes from. A **debit card** takes it directly from the cardholder's bank account. A **credit card** draws on a credit limit, and the cardholder repays the issuer later. A **prepaid card** spends money loaded onto it in advance. In Europe, most cards are debit cards.

For each payment, the issuer makes the final decision. It checks that the card is valid, that there is enough money or credit, and that the payment does not look like fraud. For online payments, it may also ask the cardholder to confirm, for example in their banking app. Once the issuer approves a payment, it collects the money from the cardholder's account or adds it to their credit card bill.

The issuer also acts for the cardholder when something goes wrong. If a payment was made without permission, or the goods never arrived, the cardholder turns to the issuer, which can claim the money back from the merchant's side through the card network.

In Europe, most cards are issued by banks, from large groups such as BNP Paribas, Santander and ING to small local banks. Since the 2010s, they have been joined by fintech companies such as Revolut and N26, which issue cards under their own licences.

## Card Networks

The card network, also called the card scheme, connects issuers and acquirers. Banks join a network as members and agree to its rules, and in return they can reach every other member. The network carries the payment messages between them, calculates each day what every bank owes the others, and sets the rules for everything from card design to disputes. Its logo on the card shows where the card is accepted.

In the four-party model, the network itself issues no cards and signs up no merchants. It leaves that to its member banks. Visa and Mastercard work this way, and together they handle most card payments in Europe. Outside Europe, other large networks include UnionPay in China and JCB in Japan.

A **three-party network** works differently: it is also the issuer and the acquirer. American Express mostly issues its own cards, signs up its own merchants and carries the payments between them, so it controls each payment from start to finish.

Several European countries also have their own **domestic card schemes**, such as Cartes Bancaires in France, Girocard in Germany and Bancomat in Italy. They do the same job as Visa and Mastercard, but only within one country, between local banks. Because they do not work abroad, their cards usually also carry a Visa or Mastercard logo, and that network handles the payment outside the country.

## Acquiring Banks

The acquiring bank, or acquirer, represents the merchant in the payment. It signs up the merchant, accepts card payments on the merchant's behalf, and pays out the money.

Before accepting a merchant, the acquirer checks who it is and what it sells, then signs a merchant agreement that sets the fees and conditions. For each payment, the acquirer sends the request through the card network to the issuer and passes the answer back. Once the issuer pays, the acquirer pays the merchant, usually within a few days, minus its fees.

The acquirer also carries the risk on the merchant's side. If a cardholder disputes a payment and wins, the acquirer must return the money, and it then recovers it from the merchant. If the merchant has closed or cannot pay, the loss stays with the acquirer. Businesses the acquirer sees as higher risk may pay more or find it harder to get accepted.

Like issuers, acquirers join each card network whose cards they accept. Many of them are no longer traditional banks, but payment companies such as Worldline, Nexi and Adyen, which hold their own licences.

## Processors and Gateways

Between the merchant's website and the card network, the payment passes through a technical layer that most customers never see. It is made up of gateways and processors.

The **payment gateway** connects the merchant's website to the payment system. At checkout, it collects the card details, encrypts them and sends them on for approval, then passes the answer back to the website. It does online what a card terminal does in a shop.

The **processor** runs the payment systems of a bank. For an acquirer, it sends the merchant's payments to the card network and works out how much each merchant is owed. For an issuer, it receives payments made with the issuer's cards and approves or declines them, following rules the issuer has set. Many banks hire a processor instead of running these systems themselves. Large processors include Fiserv, Worldline and Nexi.

Several of these roles are often held by one company. Worldline and Nexi act as both processor and acquirer. Adyen is a gateway, processor and acquirer.

## Payment Facilitators

A **payment facilitator** signs up merchants under its own account with an acquirer. The acquirer holds a single agreement with the facilitator instead of a separate one with each merchant.

The facilitator does the work an acquirer would do for each of these merchants. It checks the business, accepts its payments and pays out its money. The facilitator also carries the merchant's risk, paying the acquirer when a cardholder wins a dispute and the merchant cannot repay the money.

This opens card payments to small businesses. A merchant signs up online with the facilitator and pays a flat fee per payment. Stripe and PayPal work this way.

Companies that offer merchants a complete service for accepting payments are commonly called **payment service providers**, or PSPs. The term is broad: a PSP may be a facilitator such as Stripe, or an acquirer with its own gateway such as Adyen.

## Digital Wallets

A **digital wallet** stores the customer's cards on their phone or in their browser and lets them pay without typing card details. Apple Pay and Google Pay are the best known.

The wallet is not a new party in the four-party model. When a card is added, the issuer approves it, and the card network creates a token to replace the card number, which the wallet stores. At checkout, the wallet confirms the customer with a fingerprint, a face scan or a code, then sends the token to the merchant's gateway. From there, the payment follows the usual path through the acquirer and the card network to the issuer.

The wallet's contribution is the step at the start. It replaces typing the card details with a single confirmation, and it gives the issuer stronger proof that the real cardholder is paying. For the merchant, the result is an ordinary card payment, paid out by the same acquirer.

Some wallets go further. PayPal, for example, can also hold a balance of its own and pay the merchant from it, without using a card. Payments made this way are not card payments.

## Division of Roles

A card system can be run by one company. American Express comes close: it issues the cards, signs up the merchants and runs the network itself. Most card payments, however, pass through several companies, each doing one part of the work.

The first reason is reach. A single bank can only sign up so many cardholders and merchants. Through a shared network, its cards work at every merchant signed up by the other member banks, and its merchants accept cards from all of them. Instead of reaching an agreement with every other bank, each bank only needs to join the network.

The second reason is risk. The issuer holds the cardholder's money or lends to them, so it knows best whether the cardholder can pay. The acquirer stands behind the merchant, so it knows best whether the merchant will deliver. Splitting the roles leaves the cardholder's risk with the issuer and the merchant's risk with the acquirer.

The third reason is scale. The systems that move card payments work the same way for every bank, cost a great deal to build and must run without interruption. A processor builds them once and runs them for hundreds of banks, sharing the cost among all of them.

The result is a chain in which no single company controls the whole payment, and each one relies on the others to complete it.

## Market Power

Visa and Mastercard hold the most power. Their networks reach most cardholders in Europe. A bank that does not issue their cards, or a merchant that does not accept them, loses a large share of its business. Banks and merchants therefore accept the networks' rules and fees.

Most countries have several acquirers and processors, so banks and merchants can move from one to another. Their terms are therefore open to negotiation. Merchants with a larger volume of payments negotiate from a stronger position.

Gateways and payment facilitators are easy to sign up with, offer a simple service and charge a small fee per payment. Merchants can choose the one that suits them best, and change it when it no longer does.

## Fees

Most of the money in a card payment comes from one fee, the **merchant service charge**, which the merchant pays on every payment. The cardholder usually pays nothing for the payment itself. The charge is divided among the players along the chain.

The issuer receives the **interchange fee**. It is set by the card network and paid by the acquirer to the issuer on every payment. In the EU, it is capped at 0.2% of the payment for consumer debit cards and 0.3% for consumer credit cards.

The card network charges a **scheme fee** to both the issuer and the acquirer for the use of its network. What remains of the merchant service charge is the **acquirer's markup**, which covers the acquirer's costs and profit. When a merchant uses a payment facilitator, all of these are combined into one fee per payment.

Some players also earn outside the payment. Issuers charge cardholders for their cards and earn interest on unpaid credit card balances. Processors and gateways charge the banks and merchants they work for, per payment or per month.

For an online payment of €100 with a European consumer debit card, made through a payment facilitator, the merchant pays around €1.50 to €2. Up to €0.20 goes to the issuer, around €0.15 to the card network, and the rest to the facilitator and the companies it works with.

## Croatia

Croatia follows the same model as the rest of Europe, and most of its players belong to larger European groups. The largest banks issue most of the country's cards and accept payments for most of its merchants. Most of them are owned by foreign banking groups: Zagrebačka banka by UniCredit, Privredna banka Zagreb by Intesa Sanpaolo, Erste and Raiffeisen by Austrian groups, and OTP banka by Hungary's OTP.

Most card payments run through Visa and Mastercard, while Erste & Steiermärkische Bank also issues Diners Club cards. Croatia has no domestic card scheme.

For online payments, merchants can use Croatia-based gateways such as Monri and CorvusPay. International payment facilitators such as Stripe and PayPal also serve Croatian merchants, and Croatian cardholders also use cards from fintech issuers such as Revolut.

As a member of the EU since 2013 and of the euro area since 2023, Croatia follows the same rules as the rest of the EU, including the caps on interchange fees.
