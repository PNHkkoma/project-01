package mysql

const GetSessionDataQuery = `SELECT SessionData FROM sessiondata WHERE SessionID = ?`
const InsertSessionDataQuery = `INSERT INTO sessiondata (SessionID, SessionData) VALUES (?, ?)`
