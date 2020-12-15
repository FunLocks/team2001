package team01.yowayowa.sekibutton

import android.os.Bundle
import androidx.appcompat.app.AppCompatActivity
import android.util.Log;

class SettingsActivity:AppCompatActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_settings)

        // 設定用の Fragment を表示
        supportFragmentManager
            .beginTransaction()
            .replace(R.id.settingsContainer, PreferenceFragment())
            .commit()
    }
}