package users_test

// func TestServer_HandleUsersCreate(t *testing.T) {
// 	s := newServer(teststore.New(), sessions.NewCookieStore([]byte("secret")))
// 	testCases := []struct {
// 		desc         string
// 		payload      interface{}
// 		expectedCode int
// 	}{
// 		{
// 			desc: "when params valid",
// 			payload: map[string]string{
// 				"email":    "user@example.com",
// 				"password": "password",
// 			},
// 			expectedCode: http.StatusCreated,
// 		},
// 		{
// 			desc:         "when invalid request payload",
// 			payload:      "invalid",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			desc: "when invalid params",
// 			payload: map[string]string{
// 				"email":    "invalid",
// 				"password": "1",
// 			},
// 			expectedCode: http.StatusUnprocessableEntity,
// 		},
// 	}
// 	for _, tC := range testCases {
// 		t.Run(tC.desc, func(t *testing.T) {
// 			rec := httptest.NewRecorder()
// 			b := &bytes.Buffer{}
// 			json.NewEncoder(b).Encode(tC.payload)
// 			req, _ := http.NewRequest(http.MethodPost, "/users", b)
// 			s.ServeHTTP(rec, req)
// 			assert.Equal(t, tC.expectedCode, rec.Code)
// 		})
// 	}
// }
