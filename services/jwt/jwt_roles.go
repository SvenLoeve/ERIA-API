package jwt

const RoleUser = 1
const RoleEmployee = 2
const RoleAdmin = 3

const RoleAmbulance = 4
const RoleDoctor = 5

func hasPermissions(audience string, requiredRole int, uidMatch bool) bool {

	switch requiredRole {
	case RoleUser: // User can only edit his/her own resource.
		if uidMatch || audience == "employee" || audience == "admin" {
			return true
		}
	case RoleEmployee:
		if audience == "employee" || audience == "admin" {
			return true
		}
		break
	case RoleAdmin:
		if audience == "admin" {
			return true
		}
		break
	case RoleAmbulance:
		if audience == "ambulance" || audience == "doctor" || audience == "admin" {
			return true
		}
	case RoleDoctor:
		if audience == "doctor" || audience == "admin" {
			return true
		}
	default:
		return false
	}
	return false
}
