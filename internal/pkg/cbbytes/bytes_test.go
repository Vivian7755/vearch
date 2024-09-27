// Copyright 2019 The Vearch Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.

package cbbytes

import (
	"testing"
)

func TestFloatArray(t *testing.T) {
	fa := []float32{
		0.061978027, 0, 0.0024322208, 0.011579364, 0.02539876, 0.039685875, 0.025061483, 0, 0.011964106, 0.22668292, 0, 0.04297053, 0.018453594, 0, 0.037177425, 0.02480991, 0.007826126, 0.051577665, 0.023425028, 0.011475582, 0.0019095888, 0.013776814, 0.014983976, 0.03506184, 0, 0.04430534, 0.06460679, 0.05898117, 0, 0.024282292, 0.016249334, 0.023912193, 0.09299682, 0, 0.018714048, 0.06761856, 0.025809111, 0.056197416, 0, 0.02590032, 0.018381352, 0.025648613, 0.027967073, 0, 0.06526581, 0.04139207, 0.011967562, 0.040295754, 0.045565423, 0.019792458, 0.014155187, 0, 0.048981223, 0.078069225, 0.04663591, 0.019988984, 0.03733461, 0.008756086, 0, 0.014807857, 0.07096894, 0.07858895, 0.1279875, 0.004679786, 0.024250684, 0, 0, 0.017074037, 0, 0.002629552, 0, 0.061264943, 0.0663597, 0.030265247, 0, 0.04684891, 0.055044994, 0, 0.0011162321, 0.026748057, 0, 0.02979787, 0.049727023, 0.035215493, 0.048384063, 0.072164856, 0.016016314, 0, 0.03044394, 0.017166303, 0.060566135, 0.008639058, 0.003404264, 0, 0.055610497, 0.05342417, 0.025109932, 0.037369948, 0.0028900534, 0.0396799, 0.05425708, 0, 0.006539372, 0.0014219607, 0.018216431, 0.06387678, 0.00046353412, 0.011211042, 0.10549234, 0.07960015, 0.038522918, 0, 0.037689116, 0.035982866, 0.024110256, 0.0029472294, 0.02161264, 0.036843825, 0.015701037, 0.016980404, 0.0027704213, 0.03550382, 0, 0.067925245, 0, 0.076061904, 0.004549059, 0.04119647, 0.024341123, 0.022573616, 0.036942672, 0.03398305, 0.0242569, 0.007944243, 0.026859522, 0, 0.060345773, 0.036465496, 0.0022215715, 0, 0, 0, 0.013343078, 0.028869862, 0, 0, 0, 0.1373698, 0, 0.22396517, 0, 0.050821092, 0.0015576517, 0.0039518294, 0.038914073, 0.041303758, 0.0064045996, 0.026304115, 0, 0, 0.05650068, 0.03328706, 0, 0.037762497, 0.021793982, 0.061786786, 0.03561853, 0.05563552, 0.05866452, 0.015187055, 0.0052850423, 0.007902224, 0.05185114, 0.026978185, 0.0693349, 0.010765718, 0.010029216, 0, 0.000500787, 0.030698024, 0.0070119044, 0, 0.009976505, 0.0017787169, 0.061534386, 0.039304554, 0.003785642, 0, 0.006289713, 0, 0.045439657, 0.021738837, 0, 0.025746042, 0.0109565975, 0, 0.027416281, 0, 0, 0.040431775, 0.028270742, 0.013304096, 0.011132284, 0, 0.03723401, 0.013914178, 0.104510404, 0, 0, 0.03350801, 0.025481109, 0, 0.0063368534, 0.035471212, 0, 0.032758985, 0, 0.0036525314, 0, 0.04633592, 0, 0.023038857, 0.01670609, 0.089367166, 0.0039313035, 0.025328929, 0.034061644, 0.055928532, 0.0913302, 0.04787751, 0.03978539, 0.02132153, 0.013377509, 0, 0, 0.02406035, 0.020859528, 0.02109541, 0, 0, 0.08698767, 0.03419854, 0, 0.053695403, 0.018948961, 0.051998634, 0.00041216062, 0.15261713, 0.06787608, 0.037266616, 0.021575592, 0, 0, 0.027646538, 0, 0.046549257, 0.14935963, 0, 0.031532463, 0.026455086, 0.074673265, 0.019697012, 0.009390343, 0.008292636, 0, 0, 0.037840586, 0.07004148, 0, 0.021723269, 0, 0.023197668, 0.009895433, 0, 0.0003567019, 0, 0.045447327, 0.06099353, 0, 0.006713352, 0, 0.0069246623, 0, 0.029542867, 0.019084154, 0, 0, 0.044027753, 0.15493092, 0.07655923, 0, 0.12292403, 0.039520763, 0.12842672, 0.012103221, 0.07803762, 0.04010607, 0.03753614, 0.069991246, 0.08125684, 0.0010454362, 0.038738675, 0.0015344014, 0.020792315, 0, 0.052062314, 0, 0.027184019, 0, 0.0024863556, 0.018423248, 0.014065003, 0.023253877, 0, 0.01610332, 0.036529806, 0.0015006086, 0.08050486, 0, 0.031190122, 0.10748016, 0.027277136, 0.02079113, 0.050629016, 0, 0.005438311, 0, 0, 0.003404621, 0.009230941, 0.0018121785, 0.019612126, 0.01780753, 0.017976511, 0.011893993, 0.20906983, 0, 0, 0.020648586, 0.06436902, 0.04135358, 0.16337143, 0.007879046, 0.07855137, 0, 0, 0.106715135, 0.011456252, 0, 0, 0.067796044, 0.016482718, 0.04406853, 0, 0.014155646, 0.028609702, 0.09593144, 0.0087474985, 0, 0.13208658, 0, 0, 0.064902425, 0.03509487, 0.02185683, 0.0011197289, 0.0069919624, 0.031525593, 0.016456895, 0.012751641, 0, 0.031897705, 0.02539092, 0.04607715, 0.040976208, 0.06960661, 0, 0, 0.017952539, 0.022474745, 0.018059034, 0.030210748, 0.09745478, 0.016645519, 0, 0, 0, 0, 0.021317452, 0.029534392, 0.04813654, 0.07476152, 0.01107775, 0, 0.03512891, 0.009334718, 0, 0, 0.12861717, 0.05909355, 0.0071211844, 0.022456264, 0.06183542, 0, 0, 0.039851926, 0.026381161, 0.072980136, 0.0042176484, 0.03429989, 0, 0, 0.00025676814, 0.19011389, 0.0025905194, 0, 0.04037745, 0.038096286, 0.06619528, 0, 0.021369314, 0, 0.044375923, 0, 0, 0.038601846, 0, 0, 0.0063921385, 0, 0.05058353, 0.023487527, 0.0007609059, 0, 0.008491594, 0.01844606, 0, 0.009939489, 0.020974562, 0.037654277, 0.018058483, 0.01181587, 0, 0, 0.06735272, 0.030899672, 0, 0.07306051, 0.003967327, 0.0046983077, 0.02961305, 0.010934884, 0.0780141, 0, 0.04906736, 0.057363767, 0.04098733, 0.09345143, 0, 0.04574177, 0.01354088, 0.10529311, 0.045425035, 0, 0.06558741, 0.02241738, 0.043952692, 0, 0.07343977, 0.009315588, 0.02510878, 0.024458151, 0.07239887, 0.017329002, 0.082497105, 0, 0.0048150946, 0.010111505, 0, 0.005191731, 0.023132116, 0.0011981258, 0.016305558, 0.023025457, 0, 0.07492532, 0.027895067, 0.047399946, 0.06861661, 0.07209721, 0, 0.09101196, 0.033527695, 0.004251727, 0.06302646, 0.04658206, 0.030487519, 0.014540666, 6.865195e-05, 0, 0, 0, 0.05842679, 0.024467688, 0.0010218733, 0.03878107, 0.035439696, 0.01392871, 0.03728453, 0, 0.026774736, 0.046440836,
	}

	code, err := FloatArrayByte(fa)

	if err != nil {
		t.Fatal(err)
	}

	f := ArrayByteFloat(code)

	for inde := range fa {
		if f[inde] != fa[inde] {
			t.Fatal("diff value ", f)
		}
	}
}

func TestReadFloatArray(t *testing.T) {
	fa := []float32{
		0.061978027, 0, 0.0024322208, 0.011579364, 0.02539876, 0.039685875, 0.025061483, 0, 0.011964106, 0.22668292, 0, 0.04297053, 0.018453594, 0, 0.037177425, 0.02480991, 0.007826126, 0.051577665, 0.023425028, 0.011475582, 0.0019095888, 0.013776814, 0.014983976, 0.03506184, 0, 0.04430534, 0.06460679, 0.05898117, 0, 0.024282292, 0.016249334, 0.023912193, 0.09299682, 0, 0.018714048, 0.06761856, 0.025809111, 0.056197416, 0, 0.02590032, 0.018381352, 0.025648613, 0.027967073, 0, 0.06526581, 0.04139207, 0.011967562, 0.040295754, 0.045565423, 0.019792458, 0.014155187, 0, 0.048981223, 0.078069225, 0.04663591, 0.019988984, 0.03733461, 0.008756086, 0, 0.014807857, 0.07096894, 0.07858895, 0.1279875, 0.004679786, 0.024250684, 0, 0, 0.017074037, 0, 0.002629552, 0, 0.061264943, 0.0663597, 0.030265247, 0, 0.04684891, 0.055044994, 0, 0.0011162321, 0.026748057, 0, 0.02979787, 0.049727023, 0.035215493, 0.048384063, 0.072164856, 0.016016314, 0, 0.03044394, 0.017166303, 0.060566135, 0.008639058, 0.003404264, 0, 0.055610497, 0.05342417, 0.025109932, 0.037369948, 0.0028900534, 0.0396799, 0.05425708, 0, 0.006539372, 0.0014219607, 0.018216431, 0.06387678, 0.00046353412, 0.011211042, 0.10549234, 0.07960015, 0.038522918, 0, 0.037689116, 0.035982866, 0.024110256, 0.0029472294, 0.02161264, 0.036843825, 0.015701037, 0.016980404, 0.0027704213, 0.03550382, 0, 0.067925245, 0, 0.076061904, 0.004549059, 0.04119647, 0.024341123, 0.022573616, 0.036942672, 0.03398305, 0.0242569, 0.007944243, 0.026859522, 0, 0.060345773, 0.036465496, 0.0022215715, 0, 0, 0, 0.013343078, 0.028869862, 0, 0, 0, 0.1373698, 0, 0.22396517, 0, 0.050821092, 0.0015576517, 0.0039518294, 0.038914073, 0.041303758, 0.0064045996, 0.026304115, 0, 0, 0.05650068, 0.03328706, 0, 0.037762497, 0.021793982, 0.061786786, 0.03561853, 0.05563552, 0.05866452, 0.015187055, 0.0052850423, 0.007902224, 0.05185114, 0.026978185, 0.0693349, 0.010765718, 0.010029216, 0, 0.000500787, 0.030698024, 0.0070119044, 0, 0.009976505, 0.0017787169, 0.061534386, 0.039304554, 0.003785642, 0, 0.006289713, 0, 0.045439657, 0.021738837, 0, 0.025746042, 0.0109565975, 0, 0.027416281, 0, 0, 0.040431775, 0.028270742, 0.013304096, 0.011132284, 0, 0.03723401, 0.013914178, 0.104510404, 0, 0, 0.03350801, 0.025481109, 0, 0.0063368534, 0.035471212, 0, 0.032758985, 0, 0.0036525314, 0, 0.04633592, 0, 0.023038857, 0.01670609, 0.089367166, 0.0039313035, 0.025328929, 0.034061644, 0.055928532, 0.0913302, 0.04787751, 0.03978539, 0.02132153, 0.013377509, 0, 0, 0.02406035, 0.020859528, 0.02109541, 0, 0, 0.08698767, 0.03419854, 0, 0.053695403, 0.018948961, 0.051998634, 0.00041216062, 0.15261713, 0.06787608, 0.037266616, 0.021575592, 0, 0, 0.027646538, 0, 0.046549257, 0.14935963, 0, 0.031532463, 0.026455086, 0.074673265, 0.019697012, 0.009390343, 0.008292636, 0, 0, 0.037840586, 0.07004148, 0, 0.021723269, 0, 0.023197668, 0.009895433, 0, 0.0003567019, 0, 0.045447327, 0.06099353, 0, 0.006713352, 0, 0.0069246623, 0, 0.029542867, 0.019084154, 0, 0, 0.044027753, 0.15493092, 0.07655923, 0, 0.12292403, 0.039520763, 0.12842672, 0.012103221, 0.07803762, 0.04010607, 0.03753614, 0.069991246, 0.08125684, 0.0010454362, 0.038738675, 0.0015344014, 0.020792315, 0, 0.052062314, 0, 0.027184019, 0, 0.0024863556, 0.018423248, 0.014065003, 0.023253877, 0, 0.01610332, 0.036529806, 0.0015006086, 0.08050486, 0, 0.031190122, 0.10748016, 0.027277136, 0.02079113, 0.050629016, 0, 0.005438311, 0, 0, 0.003404621, 0.009230941, 0.0018121785, 0.019612126, 0.01780753, 0.017976511, 0.011893993, 0.20906983, 0, 0, 0.020648586, 0.06436902, 0.04135358, 0.16337143, 0.007879046, 0.07855137, 0, 0, 0.106715135, 0.011456252, 0, 0, 0.067796044, 0.016482718, 0.04406853, 0, 0.014155646, 0.028609702, 0.09593144, 0.0087474985, 0, 0.13208658, 0, 0, 0.064902425, 0.03509487, 0.02185683, 0.0011197289, 0.0069919624, 0.031525593, 0.016456895, 0.012751641, 0, 0.031897705, 0.02539092, 0.04607715, 0.040976208, 0.06960661, 0, 0, 0.017952539, 0.022474745, 0.018059034, 0.030210748, 0.09745478, 0.016645519, 0, 0, 0, 0, 0.021317452, 0.029534392, 0.04813654, 0.07476152, 0.01107775, 0, 0.03512891, 0.009334718, 0, 0, 0.12861717, 0.05909355, 0.0071211844, 0.022456264, 0.06183542, 0, 0, 0.039851926, 0.026381161, 0.072980136, 0.0042176484, 0.03429989, 0, 0, 0.00025676814, 0.19011389, 0.0025905194, 0, 0.04037745, 0.038096286, 0.06619528, 0, 0.021369314, 0, 0.044375923, 0, 0, 0.038601846, 0, 0, 0.0063921385, 0, 0.05058353, 0.023487527, 0.0007609059, 0, 0.008491594, 0.01844606, 0, 0.009939489, 0.020974562, 0.037654277, 0.018058483, 0.01181587, 0, 0, 0.06735272, 0.030899672, 0, 0.07306051, 0.003967327, 0.0046983077, 0.02961305, 0.010934884, 0.0780141, 0, 0.04906736, 0.057363767, 0.04098733, 0.09345143, 0, 0.04574177, 0.01354088, 0.10529311, 0.045425035, 0, 0.06558741, 0.02241738, 0.043952692, 0, 0.07343977, 0.009315588, 0.02510878, 0.024458151, 0.07239887, 0.017329002, 0.082497105, 0, 0.0048150946, 0.010111505, 0, 0.005191731, 0.023132116, 0.0011981258, 0.016305558, 0.023025457, 0, 0.07492532, 0.027895067, 0.047399946, 0.06861661, 0.07209721, 0, 0.09101196, 0.033527695, 0.004251727, 0.06302646, 0.04658206, 0.030487519, 0.014540666, 6.865195e-05, 0, 0, 0, 0.05842679, 0.024467688, 0.0010218733, 0.03878107, 0.035439696, 0.01392871, 0.03728453, 0, 0.026774736, 0.046440836,
	}

	code, err := FloatArrayByte(fa)

	if err != nil {
		t.Fatal(err)
	}

	f := ByteToFloat32(code)

	print(f)
}

func BenchmarkVectorToByte(b *testing.B) {
	fa := []float32{
		0.061978027, 0, 0.0024322208, 0.011579364, 0.02539876, 0.039685875, 0.025061483, 0, 0.011964106, 0.22668292, 0, 0.04297053, 0.018453594, 0, 0.037177425, 0.02480991, 0.007826126, 0.051577665, 0.023425028, 0.011475582, 0.0019095888, 0.013776814, 0.014983976, 0.03506184, 0, 0.04430534, 0.06460679, 0.05898117, 0, 0.024282292, 0.016249334, 0.023912193, 0.09299682, 0, 0.018714048, 0.06761856, 0.025809111, 0.056197416, 0, 0.02590032, 0.018381352, 0.025648613, 0.027967073, 0, 0.06526581, 0.04139207, 0.011967562, 0.040295754, 0.045565423, 0.019792458, 0.014155187, 0, 0.048981223, 0.078069225, 0.04663591, 0.019988984, 0.03733461, 0.008756086, 0, 0.014807857, 0.07096894, 0.07858895, 0.1279875, 0.004679786, 0.024250684, 0, 0, 0.017074037, 0, 0.002629552, 0, 0.061264943, 0.0663597, 0.030265247, 0, 0.04684891, 0.055044994, 0, 0.0011162321, 0.026748057, 0, 0.02979787, 0.049727023, 0.035215493, 0.048384063, 0.072164856, 0.016016314, 0, 0.03044394, 0.017166303, 0.060566135, 0.008639058, 0.003404264, 0, 0.055610497, 0.05342417, 0.025109932, 0.037369948, 0.0028900534, 0.0396799, 0.05425708, 0, 0.006539372, 0.0014219607, 0.018216431, 0.06387678, 0.00046353412, 0.011211042, 0.10549234, 0.07960015, 0.038522918, 0, 0.037689116, 0.035982866, 0.024110256, 0.0029472294, 0.02161264, 0.036843825, 0.015701037, 0.016980404, 0.0027704213, 0.03550382, 0, 0.067925245, 0, 0.076061904, 0.004549059, 0.04119647, 0.024341123, 0.022573616, 0.036942672, 0.03398305, 0.0242569, 0.007944243, 0.026859522, 0, 0.060345773, 0.036465496, 0.0022215715, 0, 0, 0, 0.013343078, 0.028869862, 0, 0, 0, 0.1373698, 0, 0.22396517, 0, 0.050821092, 0.0015576517, 0.0039518294, 0.038914073, 0.041303758, 0.0064045996, 0.026304115, 0, 0, 0.05650068, 0.03328706, 0, 0.037762497, 0.021793982, 0.061786786, 0.03561853, 0.05563552, 0.05866452, 0.015187055, 0.0052850423, 0.007902224, 0.05185114, 0.026978185, 0.0693349, 0.010765718, 0.010029216, 0, 0.000500787, 0.030698024, 0.0070119044, 0, 0.009976505, 0.0017787169, 0.061534386, 0.039304554, 0.003785642, 0, 0.006289713, 0, 0.045439657, 0.021738837, 0, 0.025746042, 0.0109565975, 0, 0.027416281, 0, 0, 0.040431775, 0.028270742, 0.013304096, 0.011132284, 0, 0.03723401, 0.013914178, 0.104510404, 0, 0, 0.03350801, 0.025481109, 0, 0.0063368534, 0.035471212, 0, 0.032758985, 0, 0.0036525314, 0, 0.04633592, 0, 0.023038857, 0.01670609, 0.089367166, 0.0039313035, 0.025328929, 0.034061644, 0.055928532, 0.0913302, 0.04787751, 0.03978539, 0.02132153, 0.013377509, 0, 0, 0.02406035, 0.020859528, 0.02109541, 0, 0, 0.08698767, 0.03419854, 0, 0.053695403, 0.018948961, 0.051998634, 0.00041216062, 0.15261713, 0.06787608, 0.037266616, 0.021575592, 0, 0, 0.027646538, 0, 0.046549257, 0.14935963, 0, 0.031532463, 0.026455086, 0.074673265, 0.019697012, 0.009390343, 0.008292636, 0, 0, 0.037840586, 0.07004148, 0, 0.021723269, 0, 0.023197668, 0.009895433, 0, 0.0003567019, 0, 0.045447327, 0.06099353, 0, 0.006713352, 0, 0.0069246623, 0, 0.029542867, 0.019084154, 0, 0, 0.044027753, 0.15493092, 0.07655923, 0, 0.12292403, 0.039520763, 0.12842672, 0.012103221, 0.07803762, 0.04010607, 0.03753614, 0.069991246, 0.08125684, 0.0010454362, 0.038738675, 0.0015344014, 0.020792315, 0, 0.052062314, 0, 0.027184019, 0, 0.0024863556, 0.018423248, 0.014065003, 0.023253877, 0, 0.01610332, 0.036529806, 0.0015006086, 0.08050486, 0, 0.031190122, 0.10748016, 0.027277136, 0.02079113, 0.050629016, 0, 0.005438311, 0, 0, 0.003404621, 0.009230941, 0.0018121785, 0.019612126, 0.01780753, 0.017976511, 0.011893993, 0.20906983, 0, 0, 0.020648586, 0.06436902, 0.04135358, 0.16337143, 0.007879046, 0.07855137, 0, 0, 0.106715135, 0.011456252, 0, 0, 0.067796044, 0.016482718, 0.04406853, 0, 0.014155646, 0.028609702, 0.09593144, 0.0087474985, 0, 0.13208658, 0, 0, 0.064902425, 0.03509487, 0.02185683, 0.0011197289, 0.0069919624, 0.031525593, 0.016456895, 0.012751641, 0, 0.031897705, 0.02539092, 0.04607715, 0.040976208, 0.06960661, 0, 0, 0.017952539, 0.022474745, 0.018059034, 0.030210748, 0.09745478, 0.016645519, 0, 0, 0, 0, 0.021317452, 0.029534392, 0.04813654, 0.07476152, 0.01107775, 0, 0.03512891, 0.009334718, 0, 0, 0.12861717, 0.05909355, 0.0071211844, 0.022456264, 0.06183542, 0, 0, 0.039851926, 0.026381161, 0.072980136, 0.0042176484, 0.03429989, 0, 0, 0.00025676814, 0.19011389, 0.0025905194, 0, 0.04037745, 0.038096286, 0.06619528, 0, 0.021369314, 0, 0.044375923, 0, 0, 0.038601846, 0, 0, 0.0063921385, 0, 0.05058353, 0.023487527, 0.0007609059, 0, 0.008491594, 0.01844606, 0, 0.009939489, 0.020974562, 0.037654277, 0.018058483, 0.01181587, 0, 0, 0.06735272, 0.030899672, 0, 0.07306051, 0.003967327, 0.0046983077, 0.02961305, 0.010934884, 0.0780141, 0, 0.04906736, 0.057363767, 0.04098733, 0.09345143, 0, 0.04574177, 0.01354088, 0.10529311, 0.045425035, 0, 0.06558741, 0.02241738, 0.043952692, 0, 0.07343977, 0.009315588, 0.02510878, 0.024458151, 0.07239887, 0.017329002, 0.082497105, 0, 0.0048150946, 0.010111505, 0, 0.005191731, 0.023132116, 0.0011981258, 0.016305558, 0.023025457, 0, 0.07492532, 0.027895067, 0.047399946, 0.06861661, 0.07209721, 0, 0.09101196, 0.033527695, 0.004251727, 0.06302646, 0.04658206, 0.030487519, 0.014540666, 6.865195e-05, 0, 0, 0, 0.05842679, 0.024467688, 0.0010218733, 0.03878107, 0.035439696, 0.01392871, 0.03728453, 0, 0.026774736, 0.046440836,
	}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			VectorToByte(fa)
		}
	})
}
