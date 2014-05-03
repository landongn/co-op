App.SignupFormComponent = Ember.Component.extend({
	password: '',
	username: '',
	email: '',

	focusIn: function (e) {
		$(e.target).siblings('span').addClass('show-tooltip');
	},
	focusOut: function (e) {
		$(e.target).siblings('span').removeClass('show-tooltip');
	},

	isValid: function () {

	}.property('password', 'email', 'username')
});