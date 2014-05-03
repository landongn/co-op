App.SignupFormComponent = Ember.Component.extend({
	password: '',
	username: '',
	email: '',

	focusIn: function (e) {
		$(e.target).siblings('span.tooltip').addClass('show-tooltip');
	},
	focusOut: function (e) {
		$(e.target).siblings('span.tooltip').removeClass('show-tooltip');
	},

	isValid: function () {
		if (this.get('username').length > 3 &&
			this.get('password').length > 4 &&
			this.get('email').length > 3 || this.get('email').match(/.+@.+\..+/i) !== null) {
			return true;
		} else {
			if (this.get('username').length === 0 || this.get('username').length < 3) {
				this.set('invalidUsername', true);
				this.set('failureUsername', 'Usernames should be at least 3 letters long. Try another one.');
			}

			if (this.get('password').length === 0 || this.get('password').length < 3) {
				this.set('invalidPassword', true);
				this.set('failurePassword', 'Passwords need to have at least 3 characters. Get creative!');
			}

			var em = this.get('email');
			if (em.match(/.+@.+\..+/i) === null) {
				this.set('invalidEmail', true);
				this.set('failureEmail', 'That isn\'t an email address.  C\'mon.');
			}

			return false;
		}
	}.property('password', 'email', 'username'),

	invalidate: function (failure) {
		switch (failure.code) {
			case 200:
				break;

			case 403:
				break;

			default:
				break;
		}
	},

	keyPress: function () {
		this.setProperties({
			invalidUsername: false,
			invalidPassword: false,
			invalidEmail: false
		});
	},
	actions: {
		attemptRegistration: function () {
			if (this.get('isValid')) {
				$.post('/account/signup', {
					username: this.get("username"),
					password: this.get('password'),
					email: this.get("email")
				}, function (resp) {
					if (resp.code === 200) {
						this.transitionTo('login');
					} else {
						this.invalidate(resp);
					}
				}.bind(this));
			}
		}
	}
});