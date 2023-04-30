// item: {
//     product_id: Number,
//     product_name: String,
//     starting_pirce: Number,
//     auction_start_date: DateTime,
//     auction_end_date:  DateTime,
//     current_bid_price: Number,
//     created_at: DateTime,
//     updayted_at: DateTime,
//     is_deleted: Number
// }

// //Adding an item
// GET     - /item/list
// POST    - /item/add
// PUT     - /item/update
// DELETE  - /item/delete/:id

// user: {
//     user_id: Number,
//     user_name: String,
//     phone: String,
//     email: String,
//     created_at: DateTime,
//     updayted_at: DateTime 
// }
// //user inform add update 
// GET     - /user/list
// POST    - /user/add  // can be unique logic of email or phone number
// PUT     - /user/update
// DELETE  - /user/delete

// item_bidding: {
//     user_id: Number,
//     product_id: Number,
//     is_booked: Number,
//     bid_amount: Number,
//     created_at: DateTime,
//     updayted_at: DateTime
// }

// POST - /bid/item 
// GET -  /bid/list?query_params["sort_by_bid_amount"]



let Node = {
    val : null,
    next: null
}
function InsertNode(elem) {
    let node = Object.create(Node)
    node.val = elem;
    return node;
}



let NodeInstance       = InsertNode(4);
NodeInstance.next      = InsertNode(5);
NodeInstance.next.next = InsertNode(6);












