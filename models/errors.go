package models

// 401
// {
//     "id": 1,
//     "result": [
//         {
//             "status": {
//                 "code": -11,
//                 "message": "No permission for the resource"
//             },
//             "url": "/pm/config/device/tftest/vdom/root/system/settings"
//         }
//     ]
// }

// 404
// {
//     "id": 1,
//     "result": [
//         {
//             "status": {
//                 "code": -3,
//                 "message": "Object does not exist"
//             },
//             "url": "/pm/config/device/tftest/vdom/root/router/static/1/"
//         }
//     ]
// }

// Device Manager Database Error Code List

// 0	OK
// 1	Unknown error
// 2	Object already exists
// 3	Object does not exist
// 4	Could not resolve reference
// 5	No such command
// 6	Invalid url
// 7	Invalid option
// 8	Invalid parameter
// 9	The command is invalid for selected url
// 10	The data is invalid for selected url
// 11	No permission for the resource
// 12	Target cannot be connected
// 13	Data is not ready
// 14	No workflow session permission
// 15	Device is not in a backup mode ADOM
// 20000	OK
// 20001	Internal error
// 20002	Invalid argument
// 20003	Initialization failed
// 20004	Not exist
// 20005	Import failed
// 20006	Checkout failed
// 20007	Reload failed
// 20008	Name already in use
// 20009	IP already in use
// 20010	Serial number already in use
// 20011	No service
// 20012	Unregistered device ignored
// 20013	Object is in use
// 20014	Not a member
// 20015	Probe failed
// 20016	Insufficient number of licenses
// 20017	Lock reqired
// 20018	Locked
// 20019	Out of range
// 20020	No permission
// 20021	Already managed
// 20022	Unsupported operation
// 20023	Unsupported device model
// 20024	Unsupported device os type
// 20025	Unsupported device version
// 20026	Unsupported device model
// 20027	Group cannot be nested more than 2 levels
// 20028	Reached limit
// 20029	Name cannot be blank
// 20030	Invalid name
// 20031	Serial number does not match device model
// 20032	Device serial number does not match database
// 20033	Command failed
// 20034	Address unknown
// 20035	Already exists
// 20036	ADOM is locked
// 20037	Device cannot be managed through HA management interface
// 20038	Probe failed: device is offline
// 20039	Probe failed: certificate mismatch
// 20040	Probe failed: encryption mismatch
// 20041	Cannot add device type to this ADOM
// 20042	Probe failed: network
// 20043	Probe failed: login error
// 20044	Please save ADOM change(s) first
// 20045	Workflow session state error
// 20046	Approval failed as there are prior sessions in queue. Please review them first.
// 20047	Update workflow data failed
// 20048	Unknown DVM error
// 20049	Workflow approval matrix required
// 20050	Duplicate workflow action
// 20051	Workflow session limit reached
// 20052	Unable to perform the requested action. The session might have been discarded or removed.
// 20053	Unknown DVM error
// 20054	Internal error for workspace action
// 20055	Workspace is locked by other user
// 20056	SQLITE error for workspace action
// 20057	Open database failed for workspace action
// 20058	invalid session for workspace action
// 20059	copy database failed for workspace action
// 20060	lock package failed for workspace action
// 20061	Invalid data for workspace action
// 20062	Number of the script for the selected type has reached the max. No more script is allowed for this type
// 20063	Unable to delete because ADOM is not empty
// 20064	Unable to delete because the adom is in global assignment list or is referenced by a defined admin account
// 20065	Unable to add FAZ because the adom mode is advanced
// 20066	Unable to add FAZ because FAZ status is enabled
// 20067	Pre-shared key already exists
// 20068	Unable to delete device group because it is used by admin
// 20069	Unknown DVM error
// 20070	Cannot delete a default meta field
