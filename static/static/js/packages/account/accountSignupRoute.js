App.AccountSignupRoute = Ember.Route.extend({
	model: function () {
		return {};
	},

	renderTemplate: function () {
		this.render('account/signup');
	}
});