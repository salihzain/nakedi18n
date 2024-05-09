package accounting

import "testing"

func TestCreateAccountEn(t *testing.T) {
	expectedError := "Invalid username: ''. The username must be a nonempty string containing no spaces"
	_, err := CreateAccount("", "", "en")
	if err == nil || err.Error() != expectedError {
		t.Fatalf("Wanted: '%s', got: '%s'", expectedError, err)
	}

	expectedError = "Invalid username: 'John Doe'. The username must be a nonempty string containing no spaces"
	_, err = CreateAccount("John Doe", "", "en")
	if err == nil || err.Error() != expectedError {
		t.Fatalf("Wanted: '%s', got: '%s'", expectedError, err)
	}

	expectedError = "Invalid empty name. The name must be a nonempty string"
	_, err = CreateAccount("john_doe", "", "en")
	if err == nil || err.Error() != expectedError {
		t.Fatalf("Wanted: '%s', got: '%s'", expectedError, err)
	}

	_, err = CreateAccount("john_doe", "John", "en")
	if err != nil {
		t.Fatalf("Wanted: 'nil', got: '%s'", err)
	}
}

func TestCreateAccountAr(t *testing.T) {
	lang := "ar"

	expectedError := "خطأ في اسم المستخدم: ''. يجب ان يكون اسم المستخدم خاليا من الفراغات"
	_, err := CreateAccount("", "", lang)
	if err == nil || err.Error() != expectedError {
		t.Fatalf("Wanted: '%s', got: '%s'", expectedError, err)
	}

	expectedError = "خطأ في اسم المستخدم: 'John Doe'. يجب ان يكون اسم المستخدم خاليا من الفراغات"
	_, err = CreateAccount("John Doe", "", lang)
	if err == nil || err.Error() != expectedError {
		t.Fatalf("Wanted: '%s', got: '%s'", expectedError, err)
	}

	expectedError = "خطأ في الاسم. يرجى تسجيل اسم"
	_, err = CreateAccount("john_doe", "", lang)
	if err == nil || err.Error() != expectedError {
		t.Fatalf("Wanted: '%s', got: '%s'", expectedError, err)
	}

	_, err = CreateAccount("john_doe", "John", lang)
	if err != nil {
		t.Fatalf("Wanted: 'nil', got: '%s'", err)
	}
}
