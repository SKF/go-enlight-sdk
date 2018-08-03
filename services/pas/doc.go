// Package pas provides the client for making API
// requests to the Enlight Point Alarm Status Service.
//
// Examples
//
// Out of Window Overall Alarm
//     40 <= x      => DANGER
//     20 <= x < 40 => ALERT
//     2 < x < 20   => GOOD
//     1 < x <= 2   => ALERT
//     x <= 1       => DANGER
//
//     SetPointAlarmThresholdInput {
//         node_id: "<node_id>",
//         user_id: "<user_id>",
//         type: OVERALL_OUT_OF_WINDOW,
//         intervals: [
//             {
//                 left_bound: 40,
//                 type: LEFT_BOUNDED_RIGHT_UNBOUNDED_LEFT_CLOSED,
//                 status: DANGER
//             },
//             {
//                 left_bound: 20,
//                 right_bound: 40,
//                 type: BOUNDED_LEFT_CLOSED_RIGHT_OPEN,
//                 status: ALERT
//             },
//             {
//                 left_bound: 2,
//                 right_bound: 20,
//                 type: BOUNDED_OPEN,
//                 status: GOOD
//             },
//             {
//                 left_bound: 1,
//                 right_bound: 2,
//                 type: BOUNDED_LEFT_OPEN_RIGHT_CLOSED,
//                 status: ALERT
//             },
//             {
//                 right_bound: 1,
//                 type: LEFT_UNBOUNDED_RIGHT_BOUNDED_RIGHT_CLOSED,
//                 status: DANGER
//             }
//         ]
//     }
//
package pas
