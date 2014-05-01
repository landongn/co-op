/*jshint node:true */
var dbm = require('db-migrate');
var type = dbm.dataType;
var async = require('async');


exports.up = function (db, callback) {

	async.series([
		db.createTable('profile', {
			id: { type: 'int', primaryKey: true, autoIncrement: true },
			username: { type: 'string', unique: true, index: true },
			password: { type: 'string' },
			email: { type: 'string', unique: true, index: true },
		}),
		callback
	]);
};

exports.down = function (db, callback) {
	async.series([
		db.dropTable('profile', callback)
	]);
};
